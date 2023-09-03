use crate::protocol::common::addr::IpAddress;
use crate::protocol::trojan::{self, CRLF};
use crate::{
    protocol::common::request::InboundRequest,
    proxy::base::SupportedProtocols,
    transport::{grpc_stream::GrpcDataReaderStream, grpc_transport::Hunk},
};

use bytes::BufMut;
use once_cell::sync::OnceCell;
use std::io::{self, Error, ErrorKind};
use std::net::{IpAddr, SocketAddr};
use tokio::io::{AsyncRead, AsyncReadExt};
use tokio::net::{TcpStream, UdpSocket};
use tokio::sync::mpsc::Sender;
use tonic::Status;

const BUFFER_SIZE: usize = 4096;

/// Static life time TCP server outbound traffic handler to avoid ARC
/// The handler is initialized through init() function
static GRPC_HANDLER: OnceCell<GrpcHandler> = OnceCell::new();

/// GrpcHandler is responsible for handling outbound traffic for GRPC inbound streams
pub struct GrpcHandler {
    protocol: SupportedProtocols,
}

impl GrpcHandler {
    pub fn new() -> &'static GrpcHandler {
        GRPC_HANDLER.get_or_init(|| Self {
            protocol: SupportedProtocols::TROJAN,
        })
    }

    pub async fn handle_hunk(
        &self,
        mut client_reader: GrpcDataReaderStream<Hunk>,
        client_writer: Sender<Result<Hunk, Status>>,
        request: InboundRequest,
    ) -> io::Result<()> {
        match self.protocol {
            SupportedProtocols::TROJAN => {
                return match request.command {
                    crate::protocol::common::command::Command::Connect => {
                        let ip_port: SocketAddr = match request.addr_port.into() {
                            Ok(sa) => sa,
                            Err(e) => return Err(e),
                        };

                        // Establish connection to remote server as specified by proxy request
                        let (mut server_reader, mut server_writer) =
                            match TcpStream::connect(ip_port).await {
                                Ok(stream) => tokio::io::split(stream),
                                Err(e) => return Err(e),
                            };

                        tokio::select!(
                            _ = tokio::io::copy(&mut client_reader, &mut server_writer) => (),
                            _ = copy_server_reader_to_client_grpc_writer(&mut server_reader, client_writer) => (),
                        );

                        Ok(())
                    }
                    crate::protocol::common::command::Command::Udp => {
                        // Establish UDP connection to remote host
                        let socket = UdpSocket::bind("0.0.0.0:0").await?;

                        tokio::select!(
                            _ = trojan::packet::copy_client_reader_to_udp_socket(client_reader, &socket) => (),
                            _ = copy_udp_socket_to_client_grpc_writer(&socket, client_writer, request) => ()
                        );

                        Ok(())
                    }
                    // Trojan only supports 2 types of commands unlike SOCKS
                    // * CONNECT X'01'
                    // * UDP ASSOCIATE X'03'
                    crate::protocol::common::command::Command::Bind => Err(Error::new(
                        ErrorKind::Unsupported,
                        "Bind command is not supported in Trojan",
                    )),
                };
            }
            _ => {
                return Err(Error::new(
                    ErrorKind::Unsupported,
                    "GrpcHandler only supports Trojan for now",
                ))
            }
        }
    }
}

async fn copy_server_reader_to_client_grpc_writer<R: AsyncRead + Unpin>(
    mut reader: R,
    writer: Sender<Result<Hunk, Status>>,
) -> io::Result<()> {
    loop {
        let mut buf = Vec::with_capacity(BUFFER_SIZE);

        match reader.read_buf(&mut buf).await {
            Ok(_) => (),
            Err(_) => {
                return Err(io::Error::new(
                    io::ErrorKind::BrokenPipe,
                    "Failed to read data from remote server",
                ))
            }
        }

        match writer.send(Ok(Hunk { data: buf })).await {
            Ok(_) => (),
            Err(_) => {
                return Err(io::Error::new(
                    io::ErrorKind::BrokenPipe,
                    "Failed to send data to client",
                ))
            }
        }
    }
}

async fn copy_udp_socket_to_client_grpc_writer(
    udp_socket: &UdpSocket,
    client_sender: Sender<Result<Hunk, Status>>,
    request: InboundRequest,
) -> io::Result<()> {
    let mut udp_buffer = vec![0u8; BUFFER_SIZE];

    loop {
        let n = match udp_socket.recv(&mut udp_buffer).await {
            Ok(n) => n,
            Err(_) => {
                return Err(io::Error::new(
                    io::ErrorKind::BrokenPipe,
                    "Failed to read data from remote server",
                ))
            }
        };

        let mut buf = Vec::with_capacity(BUFFER_SIZE);

        // Write address type to remote
        buf.put_u8(request.atype as u8);

        // Write back the address of the trojan request
        match request.addr_port.ip {
            IpAddress::IpAddr(IpAddr::V4(addr)) => {
                buf.put_slice(&addr.octets());
            }
            IpAddress::IpAddr(IpAddr::V6(addr)) => {
                buf.put_slice(&addr.octets());
            }
            IpAddress::Domain(ref domain) => {
                buf.put_u8(domain.as_bytes().len() as u8);
                buf.put_slice(domain.as_bytes());
            }
        }

        // Write port, payload size, CRLF, and the payload data into the stream
        buf.put_u16(request.addr_port.port);
        buf.put_u16(n as u16);
        buf.put_u16(CRLF);
        buf.put_slice(&udp_buffer[..n]);

        match client_sender.send(Ok(Hunk { data: buf })).await {
            Ok(_) => (),
            Err(_) => {
                return Err(io::Error::new(
                    io::ErrorKind::BrokenPipe,
                    "Failed to send data to client",
                ))
            }
        }
    }
}
