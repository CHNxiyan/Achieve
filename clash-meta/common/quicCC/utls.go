package quicCC

import (
	"github.com/metacubex/quic-go"
	c "github.com/metacubex/quic-go/congestion"
)

const (
	DefaultStreamReceiveWindow     = 15728640 // 15 MB/s
	DefaultConnectionReceiveWindow = 67108864 // 64 MB/s
)

func SetCongestionController(quicConn quic.Connection, cc string, cwnd int) {
	CWND := c.ByteCount(cwnd)
	switch cc {
	case "cubic":
		quicConn.SetCongestionControl(
			NewCubicSender(
				DefaultClock{},
				GetInitialPacketSize(quicConn.RemoteAddr()),
				false,
				nil,
			),
		)
	case "new_reno":
		quicConn.SetCongestionControl(
			NewCubicSender(
				DefaultClock{},
				GetInitialPacketSize(quicConn.RemoteAddr()),
				true,
				nil,
			),
		)
	case "bbr":
		quicConn.SetCongestionControl(
			NewBBRSender(
				DefaultClock{},
				GetInitialPacketSize(quicConn.RemoteAddr()),
				CWND*InitialMaxDatagramSize,
				DefaultBBRMaxCongestionWindow*InitialMaxDatagramSize,
			),
		)
	}
}

func UseBrutal(conn quic.Connection, tx uint64) {
	conn.SetCongestionControl(NewBrutalSender(tx))
}
