package gost

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/go-log/log"
	"github.com/songgao/water"
)

func createTun(cfg TunConfig) (conn net.Conn, itf *net.Interface, err error) {
	ip, _, err := net.ParseCIDR(cfg.Addr)
	if err != nil {
		return
	}

	ifce, err := water.New(water.Config{
		DeviceType: water.TUN,
	})
	if err != nil {
		return
	}

	mtu := cfg.MTU
	if mtu <= 0 {
		mtu = DefaultMTU
	}

	peer := cfg.Peer
	if peer == "" {
		peer = ip.String()
	}
	cmd := fmt.Sprintf("ifconfig %s inet %s %s mtu %d up",
		ifce.Name(), cfg.Addr, peer, mtu)
	log.Log("[tun]", cmd)
	args := strings.Split(cmd, " ")
	if er := exec.Command(args[0], args[1:]...).Run(); er != nil {
		err = fmt.Errorf("%s: %v", cmd, er)
		return
	}

	if err = addTunRoutes(ifce.Name(), cfg.Routes...); err != nil {
		return
	}

	itf, err = net.InterfaceByName(ifce.Name())
	if err != nil {
		return
	}

	conn = &tunTapConn{
		ifce: ifce,
		addr: &net.IPAddr{IP: ip},
	}
	return
}

func createTap(cfg TapConfig) (conn net.Conn, itf *net.Interface, err error) {
	err = errors.New("tap is not supported on darwin")
	return
}

func addTunRoutes(ifName string, routes ...IPRoute) error {
	for _, route := range routes {
		if route.Dest == nil {
			continue
		}
		cmd := fmt.Sprintf("route add -net %s -interface %s", route.Dest.String(), ifName)
		log.Log("[tun]", cmd)
		args := strings.Split(cmd, " ")
		if er := exec.Command(args[0], args[1:]...).Run(); er != nil {
			return fmt.Errorf("%s: %v", cmd, er)
		}
	}
	return nil
}
