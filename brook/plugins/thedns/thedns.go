// Copyright (c) 2016-present Cloud <cloud@txthinking.com>
//
// This program is free software; you can redistribute it and/or
// modify it under the terms of version 3 of the GNU General Public
// License as published by the Free Software Foundation.
//
// This program is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU
// General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.

package thedns

import (
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/miekg/dns"
	"github.com/patrickmn/go-cache"
	"github.com/txthinking/brook"
)

type TheDNS struct {
	BlockDomain     map[string]byte
	BypassDomain    map[string]byte
	BypassDNSClient *brook.DNSClient
	BypassDOHClient *brook.DOHClient
	DisableA        bool
	DisableAAAA     bool
	Cache           *cache.Cache
	DOHClient       *brook.DOHClient
}

func NewTheDNS(blockDomainList, bypassDomainList, bypassDNS string, disableA, disableAAAA bool, doh string) (*TheDNS, error) {
	var err error
	var ds map[string]byte
	if blockDomainList != "" {
		ds, err = brook.ReadDomainList(blockDomainList)
		if err != nil {
			return nil, err
		}
	}
	var ds1 map[string]byte
	if bypassDomainList != "" {
		ds1, err = brook.ReadDomainList(bypassDomainList)
		if err != nil {
			return nil, err
		}
	}
	var c *brook.DNSClient
	var bdc, dc *brook.DOHClient
	if !strings.HasPrefix(bypassDNS, "https://") {
		c = &brook.DNSClient{Server: bypassDNS}
	}
	if strings.HasPrefix(bypassDNS, "https://") {
		bdc, err = brook.NewDOHClient(bypassDNS)
		if err != nil {
			return nil, err
		}
	}
	if doh != "" {
		dc, err = brook.NewDOHClient(doh)
		if err != nil {
			return nil, err
		}
	}
	b := &TheDNS{
		BlockDomain:     ds,
		BypassDomain:    ds1,
		BypassDNSClient: c,
		BypassDOHClient: bdc,
		DisableA:        disableA,
		DisableAAAA:     disableAAAA,
		Cache:           cache.New(cache.NoExpiration, cache.NoExpiration),
		DOHClient:       dc,
	}
	return b, nil
}

func soa(addr *net.UDPAddr, m *dns.Msg, l1 *net.UDPConn) error {
	m1 := &dns.Msg{}
	m1.SetReply(m)
	m1.Authoritative = true
	m1.Answer = append(m1.Answer, &dns.SOA{
		Hdr:     dns.RR_Header{Name: m.Question[0].Name, Rrtype: dns.TypeSOA, Class: dns.ClassINET, Ttl: 60},
		Ns:      "txthinking.com.",
		Mbox:    "cloud.txthinking.com.",
		Serial:  uint32((time.Now().Year() * 10000) + (int(time.Now().Month()) * 100) + (time.Now().Day())*100),
		Refresh: 21600,
		Retry:   3600,
		Expire:  259200,
		Minttl:  300,
	})
	m1b, err := m1.PackBuffer(nil)
	if err != nil {
		return err
	}
	if _, err := l1.WriteToUDP(m1b, addr); err != nil {
		return err
	}
	return nil
}

func soah(m *dns.Msg, w http.ResponseWriter) error {
	m1 := &dns.Msg{}
	m1.SetReply(m)
	m1.Authoritative = true
	m1.Answer = append(m1.Answer, &dns.SOA{
		Hdr:     dns.RR_Header{Name: m.Question[0].Name, Rrtype: dns.TypeSOA, Class: dns.ClassINET, Ttl: 60},
		Ns:      "txthinking.com.",
		Mbox:    "cloud.txthinking.com.",
		Serial:  uint32((time.Now().Year() * 10000) + (int(time.Now().Month()) * 100) + (time.Now().Day())*100),
		Refresh: 21600,
		Retry:   3600,
		Expire:  259200,
		Minttl:  300,
	})
	m1b, err := m1.PackBuffer(nil)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/dns-message")
	w.Write(m1b)
	return nil
}

func (p *TheDNS) TouchBrook() {
	f := brook.DNSGate
	brook.DNSGate = func(addr *net.UDPAddr, m *dns.Msg, l1 *net.UDPConn) (bool, error) {
		done, err := f(addr, m, l1)
		if err != nil || done {
			return done, err
		}
		if m.Question[0].Qtype == dns.TypeA && p.DisableA {
			err := soa(addr, m, l1)
			return err == nil, err
		}
		if m.Question[0].Qtype == dns.TypeAAAA && p.DisableAAAA {
			err := soa(addr, m, l1)
			return err == nil, err
		}
		if brook.ListHasDomain(p.BlockDomain, m.Question[0].Name[0:len(m.Question[0].Name)-1], p.Cache) {
			err := soa(addr, m, l1)
			return err == nil, err
		}
		if brook.ListHasDomain(p.BypassDomain, m.Question[0].Name[0:len(m.Question[0].Name)-1], p.Cache) {
			var m1 *dns.Msg
			if p.BypassDNSClient != nil {
				m1, err = p.BypassDNSClient.Exchange(m)
			}
			if p.BypassDOHClient != nil {
				m1, err = p.BypassDOHClient.Exchange(m)
			}
			if err != nil {
				return false, err
			}
			m1b, err := m1.PackBuffer(nil)
			if err != nil {
				return false, err
			}
			if _, err := l1.WriteToUDP(m1b, addr); err != nil {
				return false, err
			}
			return true, nil
		}
		if p.DOHClient != nil {
			m1, err := p.DOHClient.Exchange(m)
			if err != nil {
				return false, err
			}
			m1b, err := m1.PackBuffer(nil)
			if err != nil {
				return false, err
			}
			if _, err := l1.WriteToUDP(m1b, addr); err != nil {
				return false, err
			}
			return true, nil
		}
		return false, nil
	}
	f1 := brook.DOHGate
	brook.DOHGate = func(m *dns.Msg, w http.ResponseWriter, r *http.Request) (done bool, err error) {
		done, err = f1(m, w, r)
		if err != nil || done {
			return done, err
		}
		if m.Question[0].Qtype == dns.TypeA && p.DisableA {
			err := soah(m, w)
			return err == nil, err
		}
		if m.Question[0].Qtype == dns.TypeAAAA && p.DisableAAAA {
			err := soah(m, w)
			return err == nil, err
		}
		if brook.ListHasDomain(p.BlockDomain, m.Question[0].Name[0:len(m.Question[0].Name)-1], p.Cache) {
			err := soah(m, w)
			return err == nil, err
		}
		return false, nil
	}
}
