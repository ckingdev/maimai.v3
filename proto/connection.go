package proto

import (
	"euphoria.io/heim/proto"
)

type Connection interface {
	Receive() (*proto.Packet, error)
	Send(p *proto.Packet) error
	Close() error
}

type Dialer interface {
	Dial() (Connection, error)
}
