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
	"net"
)

type Exchanger interface {
	Network() string
	Src() string
	Dst() string
	Exchange(remote net.Conn) error
	Clean()
}

type UDPServerConnFactory interface {
	Handle(addr *net.UDPAddr, b, p []byte, w func([]byte) (int, error), timeout int) (net.Conn, []byte, error)
}

var ServerGate func(ex Exchanger) (Exchanger, error) = func(ex Exchanger) (Exchanger, error) {
	return ex, nil
}

var ClientGate func(ex Exchanger) (Exchanger, error) = func(ex Exchanger) (Exchanger, error) {
	return ex, nil
}
