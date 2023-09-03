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

package brook

import (
	"context"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	cache "github.com/patrickmn/go-cache"
	"github.com/txthinking/brook/limits"
	crypto1 "github.com/txthinking/crypto"
	"github.com/txthinking/socks5"
	"github.com/urfave/negroni"
	"golang.org/x/crypto/acme/autocert"
)

// WSServer.
type WSServer struct {
	Password       []byte
	Domain         string
	TCPAddr        *net.TCPAddr
	HTTPServer     *http.Server
	HTTPSServer    *http.Server
	TCPTimeout     int
	UDPTimeout     int
	Path           string
	UDPSrc         *cache.Cache
	BlockDomain    map[string]byte
	BlockCIDR4     []*net.IPNet
	BlockCIDR6     []*net.IPNet
	BlockCache     *cache.Cache
	BlockLock      *sync.RWMutex
	Done           chan byte
	WSSServerPort  int64
	Cert           []byte
	CertKey        []byte
	WithoutBrook   bool
	PasswordSha256 []byte
	Dial           func(network, laddr, raddr string) (net.Conn, error)
}

// NewWSServer.
func NewWSServer(addr, password, domain, path string, tcpTimeout, udpTimeout int, blockDomainList, blockCIDR4List, blockCIDR6List string, updateListInterval int64) (*WSServer, error) {
	var taddr *net.TCPAddr
	var err error
	if domain == "" {
		taddr, err = net.ResolveTCPAddr("tcp", addr)
		if err != nil {
			return nil, err
		}
	}
	var ds map[string]byte
	if blockDomainList != "" {
		ds, err = ReadDomainList(blockDomainList)
		if err != nil {
			return nil, err
		}
	}
	var c4 []*net.IPNet
	if blockCIDR4List != "" {
		c4, err = ReadCIDRList(blockCIDR4List)
		if err != nil {
			return nil, err
		}
	}
	var c6 []*net.IPNet
	if blockCIDR6List != "" {
		c6, err = ReadCIDRList(blockCIDR6List)
		if err != nil {
			return nil, err
		}
	}
	cs2 := cache.New(cache.NoExpiration, cache.NoExpiration)
	cs3 := cache.New(cache.NoExpiration, cache.NoExpiration)
	var lock *sync.RWMutex
	if updateListInterval != 0 {
		lock = &sync.RWMutex{}
	}
	done := make(chan byte)
	if err := limits.Raise(); err != nil {
		log.Println("Try to raise system limits, got", err)
	}
	b, err := crypto1.SHA256Bytes([]byte(password))
	if err != nil {
		return nil, err
	}
	s := &WSServer{
		Password:       []byte(password),
		Domain:         domain,
		TCPAddr:        taddr,
		TCPTimeout:     tcpTimeout,
		UDPTimeout:     udpTimeout,
		Path:           path,
		UDPSrc:         cs2,
		BlockDomain:    ds,
		BlockCIDR4:     c4,
		BlockCIDR6:     c6,
		BlockCache:     cs3,
		BlockLock:      lock,
		Done:           done,
		PasswordSha256: b,
	}
	if updateListInterval != 0 {
		go func() {
			ticker := time.NewTicker(time.Duration(updateListInterval) * time.Second)
			defer ticker.Stop()
			for {
				select {
				case <-done:
					return
				case <-ticker.C:
					var ds map[string]byte
					if blockDomainList != "" {
						ds, err = ReadDomainList(blockDomainList)
						if err != nil {
							log.Println("ReadDomainList", blockDomainList, err)
							break
						}
					}
					var c4 []*net.IPNet
					if blockCIDR4List != "" {
						c4, err = ReadCIDRList(blockCIDR4List)
						if err != nil {
							log.Println("ReadCIDRList", blockCIDR4List, err)
							break
						}
					}
					var c6 []*net.IPNet
					if blockCIDR6List != "" {
						c6, err = ReadCIDRList(blockCIDR6List)
						if err != nil {
							log.Println("ReadCIDRList", blockCIDR6List, err)
							break
						}
					}
					lock.Lock()
					s.BlockDomain = ds
					s.BlockCIDR4 = c4
					s.BlockCIDR6 = c6
					if cs3 != nil {
						cs3.Flush()
					}
					lock.Unlock()
				}
			}
		}()
	}
	return s, nil
}

// Run server.
func (s *WSServer) ListenAndServe() error {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		return
	})
	r.Methods("GET").Path(s.Path).Handler(s)

	n := negroni.New()
	n.Use(negroni.NewRecovery())
	if Debug {
		n.Use(negroni.NewLogger())
	}
	n.UseFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		w.Header().Set("Server", "nginx")
		next(w, r)
	})
	n.UseHandler(r)

	if s.Domain == "" {
		s.HTTPServer = &http.Server{
			Addr:           s.TCPAddr.String(),
			ReadTimeout:    5 * time.Second,
			WriteTimeout:   10 * time.Second,
			IdleTimeout:    120 * time.Second,
			MaxHeaderBytes: 1 << 20,
			Handler:        n,
		}
		return s.HTTPServer.ListenAndServe()
	}
	var t *tls.Config
	if s.Cert == nil || s.CertKey == nil {
		m := autocert.Manager{
			Cache:      autocert.DirCache(".letsencrypt"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(s.Domain),
			Email:      "cloud@txthinking.com",
		}
		go func() {
			log.Println(http.ListenAndServe(":80", m.HTTPHandler(nil)))
		}()
		t = &tls.Config{GetCertificate: m.GetCertificate}
	}
	if s.Cert != nil && s.CertKey != nil {
		ct, err := tls.X509KeyPair(s.Cert, s.CertKey)
		if err != nil {
			return err
		}
		t = &tls.Config{Certificates: []tls.Certificate{ct}, ServerName: s.Domain}
	}
	s.HTTPSServer = &http.Server{
		Addr:         ":" + strconv.FormatInt(s.WSSServerPort, 10),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      n,
		TLSConfig:    t,
	}
	go func() {
		time.Sleep(1 * time.Second)
		c := &http.Client{
			Timeout: 10 * time.Second,
		}
		_, _ = c.Get("https://" + s.Domain + s.Path)
	}()
	return s.HTTPSServer.ListenAndServeTLS("", "")
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  65507,
	WriteBufferSize: 65507,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (s *WSServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := conn.UnderlyingConn()
	defer c.Close()
	if s.TCPTimeout != 0 {
		if err := c.SetDeadline(time.Now().Add(time.Duration(s.TCPTimeout) * time.Second)); err != nil {
			log.Println(err)
			return
		}
	}
	b := s.Password
	if s.WithoutBrook {
		b = s.PasswordSha256
	}
	ss, dst, err := MakeStreamServer(b, c, s.TCPTimeout, s.WithoutBrook)
	if err != nil {
		log.Println(err)
		return
	}
	defer ss.Clean()
	if ss.NetworkName() == "tcp" {
		if err := s.TCPHandle(ss, dst); err != nil {
			log.Println(err)
		}
	}
	if ss.NetworkName() == "udp" {
		ss.SetTimeout(s.UDPTimeout)
		if err := s.UDPHandle(ss, c.RemoteAddr().String(), dst); err != nil {
			log.Println(err)
		}
	}
}

// TCPHandle handles request.
func (s *WSServer) TCPHandle(ss Exchanger, dst []byte) error {
	address := socks5.ToAddress(dst[0], dst[1:len(dst)-2], dst[len(dst)-2:])
	if Debug {
		log.Println("dial tcp", address)
	}
	var ds map[string]byte
	var c4 []*net.IPNet
	var c6 []*net.IPNet
	if s.BlockLock != nil {
		s.BlockLock.RLock()
	}
	ds = s.BlockDomain
	c4 = s.BlockCIDR4
	c6 = s.BlockCIDR6
	if s.BlockLock != nil {
		s.BlockLock.RUnlock()
	}
	if BlockAddress(address, ds, c4, c6, s.BlockCache) {
		return errors.New("block " + address)
	}
	var rc net.Conn
	var err error
	if s.Dial == nil {
		rc, err = Dial.Dial("tcp", address)
	}
	if s.Dial != nil {
		rc, err = s.Dial("tcp", "", address)
	}
	if err != nil {
		return err
	}
	defer rc.Close()
	if s.TCPTimeout != 0 {
		if err := rc.SetDeadline(time.Now().Add(time.Duration(s.TCPTimeout) * time.Second)); err != nil {
			return err
		}
	}
	if err := ss.Exchange(rc); err != nil {
		return nil
	}
	return nil
}

// UDPHandle handles packet.
func (s *WSServer) UDPHandle(ss Exchanger, src string, dstb []byte) error {
	dst := socks5.ToAddress(dstb[0], dstb[1:len(dstb)-2], dstb[len(dstb)-2:])
	if Debug {
		log.Println("dial udp", dst)
	}
	var ds map[string]byte
	var c4 []*net.IPNet
	var c6 []*net.IPNet
	if s.BlockLock != nil {
		s.BlockLock.RLock()
	}
	ds = s.BlockDomain
	c4 = s.BlockCIDR4
	c6 = s.BlockCIDR6
	if s.BlockLock != nil {
		s.BlockLock.RUnlock()
	}
	if BlockAddress(dst, ds, c4, c6, s.BlockCache) {
		return errors.New("block " + dst)
	}
	var laddr *net.UDPAddr
	any, ok := s.UDPSrc.Get(src + dst)
	if ok {
		laddr = any.(*net.UDPAddr)
	}
	raddr, err := net.ResolveUDPAddr("udp", dst)
	if err != nil {
		return err
	}
	var rc net.Conn
	if s.Dial == nil {
		rc, err = Dial.DialUDP("udp", laddr, raddr)
	}
	if s.Dial != nil {
		la := ""
		if laddr != nil {
			la = laddr.String()
		}
		rc, err = s.Dial("udp", la, dst)
	}
	if err != nil {
		return err
	}
	defer rc.Close()
	if s.UDPTimeout != 0 {
		if err := rc.SetDeadline(time.Now().Add(time.Duration(s.UDPTimeout) * time.Second)); err != nil {
			return err
		}
	}
	if laddr == nil {
		s.UDPSrc.Set(src+dst, rc.LocalAddr().(*net.UDPAddr), -1)
	}
	if err := ss.Exchange(rc); err != nil {
		return nil
	}
	return nil
}

// Shutdown server.
func (s *WSServer) Shutdown() error {
	close(s.Done)
	if s.Domain == "" {
		return s.HTTPServer.Shutdown(context.Background())
	}
	return s.HTTPSServer.Shutdown(context.Background())
}
