package v2raywebsocket

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/sagernet/sing-box/adapter"
	"github.com/sagernet/sing-box/common/tls"
	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
	sHTTP "github.com/sagernet/sing/protocol/http"

	"github.com/gobwas/ws"
)

var _ adapter.V2RayClientTransport = (*Client)(nil)

type Client struct {
	dialer              N.Dialer
	tlsConfig           tls.Config
	serverAddr          M.Socksaddr
	requestURL          url.URL
	headers             http.Header
	maxEarlyData        uint32
	earlyDataHeaderName string
}

func NewClient(ctx context.Context, dialer N.Dialer, serverAddr M.Socksaddr, options option.V2RayWebsocketOptions, tlsConfig tls.Config) adapter.V2RayClientTransport {
	if tlsConfig != nil {
		if len(tlsConfig.NextProtos()) == 0 {
			tlsConfig.SetNextProtos([]string{"http/1.1"})
		}
	}
	var requestURL url.URL
	if tlsConfig == nil {
		requestURL.Scheme = "ws"
	} else {
		requestURL.Scheme = "wss"
	}
	requestURL.Host = serverAddr.String()
	requestURL.Path = options.Path
	err := sHTTP.URLSetPath(&requestURL, options.Path)
	if err != nil {
		return nil
	}
	if !strings.HasPrefix(requestURL.Path, "/") {
		requestURL.Path = "/" + requestURL.Path
	}
	headers := make(http.Header)
	for key, value := range options.Headers {
		headers[key] = value
	}
	return &Client{
		dialer,
		tlsConfig,
		serverAddr,
		requestURL,
		headers,
		options.MaxEarlyData,
		options.EarlyDataHeaderName,
	}
}

func (c *Client) dialContext(ctx context.Context, requestURL *url.URL, headers http.Header) (*WebsocketConn, error) {
	conn, err := c.dialer.DialContext(ctx, N.NetworkTCP, c.serverAddr)
	if err != nil {
		return nil, err
	}
	if c.tlsConfig != nil {
		conn, err = tls.ClientHandshake(ctx, conn, c.tlsConfig)
		if err != nil {
			return nil, err
		}
	}
	conn.SetDeadline(time.Now().Add(C.TCPTimeout))
	reader, _, err := ws.Dialer{Header: ws.HandshakeHeaderHTTP(headers)}.Upgrade(conn, requestURL)
	conn.SetDeadline(time.Time{})
	if err != nil {
		return nil, err
	}
	return NewConn(conn, reader, nil, ws.StateClientSide), nil
}

func (c *Client) DialContext(ctx context.Context) (net.Conn, error) {
	if c.maxEarlyData <= 0 {
		conn, err := c.dialContext(ctx, &c.requestURL, c.headers)
		if err != nil {
			return nil, err
		}
		return conn, nil
	} else {
		return &EarlyWebsocketConn{Client: c, ctx: ctx, create: make(chan struct{})}, nil
	}
}
