package main

import (
	"context"
	"net"
	"net/http"
	"net/netip"
	"testing"

	C "github.com/sagernet/sing-box/constant"
	"github.com/sagernet/sing-box/option"
	"github.com/sagernet/sing-dns"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"
	"github.com/sagernet/sing/protocol/socks"
)

func TestQUICDNS(t *testing.T) {
	for _, dnsServer := range []string{
		"quic://dns.adguard-dns.com",
		"h3://1.1.1.1/dns-query",
	} {
		t.Run(dnsServer, func(t *testing.T) {
			testDNS(t, dnsServer)
		})
	}
}

func testDNS(t *testing.T, address string) {
	startInstance(t, option.Options{
		DNS: &option.DNSOptions{
			Servers: []option.DNSServerOptions{
				{
					Address:         address,
					AddressResolver: "local",
				},
				{
					Address: "local",
					Tag:     "local",
				},
			},
		},
		Inbounds: []option.Inbound{
			{
				Type: C.TypeMixed,
				Tag:  "mixed-in",
				MixedOptions: option.HTTPMixedInboundOptions{
					ListenOptions: option.ListenOptions{
						Listen:     option.NewListenAddress(netip.IPv4Unspecified()),
						ListenPort: clientPort,
						InboundOptions: option.InboundOptions{
							DomainStrategy: option.DomainStrategy(dns.DomainStrategyUseIPv4),
						},
					},
				},
			},
		},
		Outbounds: []option.Outbound{
			{
				Type: C.TypeDirect,
			},
		},
	})
	dialer := socks.NewClient(N.SystemDialer, M.ParseSocksaddrHostPort("127.0.0.1", clientPort), socks.Version5, "", "")
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return dialer.DialContext(context.Background(), network, M.ParseSocksaddr(addr))
			},
		},
	}
	resp, err := client.Get("https://captive.apple.com")
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	client.CloseIdleConnections()
}
