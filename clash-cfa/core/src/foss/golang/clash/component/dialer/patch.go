package dialer

import (
	"context"
	"net"
	"syscall"
)

type TunnelDialer func(context context.Context, network, address string) (net.Conn, error)
type SocketControl func(network, address string, conn syscall.RawConn) error

var DefaultTunnelDialer TunnelDialer
var DefaultSocketHook SocketControl

func DialTunnelContext(ctx context.Context, network, address string) (net.Conn, error) {
	if dialer := DefaultTunnelDialer; dialer != nil {
		return dialer(ctx, network, address)
	}

	return DialContext(ctx, network, address)
}

func dialContextHooked(ctx context.Context, network string, destination net.IP, port string) (net.Conn, error) {
	dialer := &net.Dialer{
		Control: DefaultSocketHook,
	}

	conn, err := dialer.DialContext(ctx, network, net.JoinHostPort(destination.String(), port))
	if err != nil {
		return nil, err
	}

	if t, ok := conn.(*net.TCPConn); ok {
		t.SetKeepAlive(false)
	}

	return conn, nil
}

func listenPacketHooked(ctx context.Context, network, address string) (net.PacketConn, error) {
	lc := &net.ListenConfig{
		Control: DefaultSocketHook,
	}

	return lc.ListenPacket(ctx, network, address)
}