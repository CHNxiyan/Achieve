use std::mem::MaybeUninit;

use bytes::BytesMut;
use futures::TryFutureExt;
use tokio::io::{AsyncReadExt, AsyncWriteExt, DuplexStream, ReadBuf};
use trojan_rust::config::base::InboundConfig;
use trojan_rust::proxy::base::SupportedProtocols;

#[test]
fn test_acceptor_initialization() {
    // let inbound_config = InboundConfig {
    //     address: "1.2.3.4".to_string(),
    //     port: 123,
    //     protocol: SupportedProtocols::TROJAN,
    //     secret: Some("123123".to_string()),
    //     tls: None,
    // };
    // let acceptor = Acceptor::new(&inbound_config);
    // assert_eq!(1, 1);
}

#[tokio::test]
async fn test_buffer() {
    // let mut buf: Vec<u8> = Vec::with_capacity(2);

    // let (mut server, mut client) = tokio::io::duplex(512);

    // server.write(&[2u8; 16]).await.unwrap();

    // client.read_exact(&mut buf).await.unwrap();

    // print!("{:?}", buf);
    use tokio::sync::oneshot;

    let (tx, rx) = oneshot::channel::<u32>();

    tx.send(1);
    // tx.send(2);

}
