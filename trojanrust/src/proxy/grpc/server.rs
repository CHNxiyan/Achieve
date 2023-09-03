use crate::config::base::{InboundConfig, OutboundConfig};
use crate::transport::grpc_transport::grpc_service_server::GrpcService;
use crate::transport::grpc_transport::grpc_service_server::GrpcServiceServer;
use crate::transport::grpc_transport::{Hunk, MultiHunk};

use futures::Stream;
use log::{info, warn};
use std::io::{self, Error, ErrorKind};
use std::net::ToSocketAddrs;
use std::pin::Pin;
use tokio::sync::mpsc;
use tonic::transport::{Identity, Server, ServerTlsConfig};
use tonic::{Request, Response, Status, Streaming};

use super::acceptor::GrpcAcceptor;
use super::handler::GrpcHandler;

// TODO: Need more discretion in detemining the value for channel size, or make it configurable
const CHANNEL_SIZE: usize = 16;

/// Start running the GRPC server
pub async fn start(
    inbound_config: &'static InboundConfig,
    _outbound_config: &'static OutboundConfig,
) -> io::Result<()> {
    // Extract the address that the server should listen on
    let address = match (inbound_config.address.as_ref(), inbound_config.port)
        .to_socket_addrs()
        .unwrap()
        .next()
    {
        Some(addr) => addr,
        None => {
            return Err(Error::new(
                ErrorKind::AddrNotAvailable,
                "incorrect address in configuration",
            ))
        }
    };

    let tls_config = match &inbound_config.tls {
        Some(cfg) => {
            let cert = tokio::fs::read(cfg.cert_path.clone()).await?;
            let key = tokio::fs::read(cfg.key_path.clone()).await?;
            Some(ServerTlsConfig::new().identity(Identity::from_pem(cert, key)))
        }
        None => None,
    };

    // Initialize and start the GRPC server to serve GRPC requests
    let mut server = match tls_config {
        Some(cfg) => Server::builder()
            .tls_config(cfg)
            .expect("Failed to build GRPC server"),
        None => Server::builder(),
    };

    return match server
        .add_service(GrpcServiceServer::new(GrpcProxyService::new(
            GrpcAcceptor::new(&inbound_config),
            GrpcHandler::new(),
        )))
        .serve(address)
        .await
    {
        Ok(_) => Ok(()),
        Err(e) => Err(Error::new(
            ErrorKind::Interrupted,
            format!("Failed to start grpc server: {}", e),
        )),
    };
}

pub struct GrpcProxyService {
    acceptor: &'static GrpcAcceptor,
    handler: &'static GrpcHandler,
}

impl GrpcProxyService {
    pub fn new(acceptor: &'static GrpcAcceptor, handler: &'static GrpcHandler) -> Self {
        Self { acceptor, handler }
    }
}

#[tonic::async_trait]
impl GrpcService for GrpcProxyService {
    type TunStream = Pin<Box<dyn Stream<Item = Result<Hunk, Status>> + Send>>;
    type TunMultiStream = Pin<Box<dyn Stream<Item = Result<MultiHunk, Status>> + Send>>;

    async fn tun(
        &self,
        request: Request<Streaming<Hunk>>,
    ) -> Result<Response<Self::TunStream>, Status> {
        info!("Received GRPC request");

        let (acceptor, handler) = (self.acceptor, self.handler);
        let (tx, rx) = mpsc::channel(CHANNEL_SIZE);

        tokio::spawn(async move {
            let (request, client_reader) = match acceptor.accept_hunk(request).await {
                Ok((req, reader)) => (req, reader),
                Err(e) => {
                    warn!("Failed to accept the inbound traffic: {}", e);
                    return;
                }
            };

            match handler.handle_hunk(client_reader, tx, request).await {
                Ok(_) => return,
                Err(e) => {
                    warn!("Failed to handle inbound traffic: {}", e);
                    return;
                }
            };
        });

        Ok(Response::new(Box::pin(
            tokio_stream::wrappers::ReceiverStream::new(rx),
        )))
    }

    async fn tun_multi(
        &self,
        _request: tonic::Request<Streaming<MultiHunk>>,
    ) -> Result<Response<Self::TunMultiStream>, Status> {
        todo!()
    }
}
