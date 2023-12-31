package mux

import (
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing-mux"
	N "github.com/sagernet/sing/common/network"
)

type Client = mux.Client

func NewClientWithOptions(dialer N.Dialer, options option.OutboundMultiplexOptions) (*Client, error) {
	if !options.Enabled {
		return nil, nil
	}
	return mux.NewClient(mux.Options{
		Dialer:         dialer,
		Protocol:       options.Protocol,
		MaxConnections: options.MaxConnections,
		MinStreams:     options.MinStreams,
		MaxStreams:     options.MaxStreams,
		Padding:        options.Padding,
	})
}
