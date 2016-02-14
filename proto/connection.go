package proto

import (
	"euphoria.io/heim/proto"
)

// Connection represents a connection to a room.
type Connection interface {

	// Receive blocks until a packet is received, and returns the packet
	// received or an error.
	Receive() (*proto.Packet, error)

	// Send sends the given packet on the Connection and returns the error, if any.
	Send(p *proto.Packet) error

	// Close ends the Connection and cleans up any necessary underlying connections.
	Close() error
}

// Dialer creates a new Connection with a configuration specified by the type
// satisfying the interface.
type Dialer interface {

	// Dial creates a new Connection, opens, any underlying connections, and returns
	// the Connection or an error.
	Dial() (Connection, error)
}
