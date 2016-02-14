package proto

import (
	"euphoria.io/heim/proto"
	"euphoria.io/scope"
)

type Handler interface {
	Run(ctx scope.Context) error
	HandleIncoming(conn Connection, p *proto.Packet) error
}
