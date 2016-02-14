package proto

import (
	"euphoria.io/heim/proto"
	"euphoria.io/scope"
)

// Handler describes an object that can run in the background and respond
// directly to incoming packets.
type Handler interface {

	// Run starts the background process for the Handler, if any.
	Run(ctx scope.Context) error

	// HandleIncoming allows the handler to respond directly to incoming packets
	// using the given Connection.
	HandleIncoming(conn Connection, p *proto.Packet) error
}
