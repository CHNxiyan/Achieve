use crate::config::base::InboundConfig;
use crate::config::tls::make_server_config;
use crate::protocol::common::request::InboundRequest;
use crate::protocol::common::stream::StandardTcpStream;
use crate::protocol::socks5;
use crate::protocol::trojan;
use crate::proxy::base::SupportedProtocols;

use once_cell::sync::OnceCell;
use sha2::{Digest, Sha224};
use std::io::{Error, ErrorKind, Result};
use std::sync::Arc;
use tokio::io::{AsyncRead, AsyncWrite};
use tokio_rustls::TlsAcceptor;

/// Static cell for storing the TCP acceptor with static lifetime to avoid use of Arc.
static TCP_ACCEPTOR: OnceCell<TcpAcceptor> = OnceCell::new();

/// Acceptor handles incomming connection by escalating them to application level data stream based on
/// the configuration. It is also responsible for escalating TCP connection to TLS connection if the user
/// enabled TLS.
pub struct TcpAcceptor {
    tls_acceptor: Option<TlsAcceptor>,
    port: u16,
    protocol: SupportedProtocols,
    secret: Vec<u8>,
}

impl TcpAcceptor {
    /// Instantiate a new acceptor based on InboundConfig passed by the user. It will generate the secret based on
    /// secret in the config file and the selected protocol and instantiate TLS acceptor is it is enabled.
    pub fn init(inbound: &InboundConfig) -> &'static Self {
        let secret = match inbound.protocol {
            SupportedProtocols::TROJAN if inbound.secret.is_some() => {
                let secret = inbound.secret.as_ref().unwrap();
                Sha224::digest(secret.as_bytes())
                    .iter()
                    .map(|x| format!("{:02x}", x))
                    .collect::<String>()
                    .as_bytes()
                    .to_vec()
            }
            _ => Vec::new(),
        };

        let tls_acceptor = match &inbound.tls {
            Some(tls) => match make_server_config(&tls) {
                Some(cfg) => Some(TlsAcceptor::from(Arc::new(cfg))),
                None => None,
            },
            None => None,
        };

        TCP_ACCEPTOR.get_or_init(|| Self {
            tls_acceptor,
            port: inbound.port,
            protocol: inbound.protocol,
            secret,
        })
    }

    /// Takes an inbound TCP stream, escalate to TLS if possible and then escalate to application level data stream
    /// to be ready to read user's request and process them.
    pub async fn accept<T: AsyncRead + AsyncWrite + Send + Unpin>(
        &self,
        inbound_stream: T,
    ) -> Result<(InboundRequest, StandardTcpStream<T>)> {
        let stream = if let Some(tls_acceptor) = self.tls_acceptor.as_ref() {
            StandardTcpStream::RustlsServer(tls_acceptor.accept(inbound_stream).await?)
        } else {
            StandardTcpStream::Plain(inbound_stream)
        };

        match self.protocol {
            // Handle SOCK5 protocol
            SupportedProtocols::SOCKS => Ok(socks5::accept(stream, self.port).await?),
            // Handle Trojan protocol
            SupportedProtocols::TROJAN => Ok(trojan::accept(stream, &self.secret).await?),
            // Shutdown the connection if the protocol is currently unsupported
            _ => Err(Error::new(
                ErrorKind::ConnectionReset,
                "Failed to accept inbound stream, unsupported protocol",
            )),
        }
    }
}
