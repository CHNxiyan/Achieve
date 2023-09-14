package interrupt

import (
	"io"
	"net"
	"sync"

	"github.com/sagernet/sing/common/x/list"
)

type Group struct {
	access      sync.Mutex
	connections list.List[*groupConnItem]
}

type groupConnItem struct {
	conn       io.Closer
	isExternal bool
}

func NewGroup() *Group {
	return &Group{}
}

func (m *Group) NewConn(conn net.Conn, isExternal bool) net.Conn {
	m.access.Lock()
	defer m.access.Unlock()
	item := m.connections.PushBack(&groupConnItem{conn, isExternal})
	return &Conn{Conn: conn, element: item}
}

func (m *Group) NewPacketConn(conn net.PacketConn, isExternal bool) net.PacketConn {
	m.access.Lock()
	defer m.access.Unlock()
	item := m.connections.PushBack(&groupConnItem{conn, isExternal})
	return &PacketConn{PacketConn: conn, element: item}
}

func (m *Group) Interrupt(interruptExternalConnections bool) {
	m.access.Lock()
	defer m.access.Unlock()
	var toDelete []*list.Element[*groupConnItem]
	for element := m.connections.Front(); element != nil; element = element.Next() {
		if !element.Value.isExternal || interruptExternalConnections {
			element.Value.conn.Close()
			toDelete = append(toDelete, element)
		}
	}
	for _, element := range toDelete {
		m.connections.Remove(element)
	}
}
