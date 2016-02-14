package connection

import (
	mproto "github.com/cpalone/maimai.v3/proto"

	"euphoria.io/heim/proto"
	"euphoria.io/scope"
)

type MockConnection struct {
	ctx      scope.Context
	incoming chan *proto.Packet
	outgoing chan *proto.Packet
}

type MockDialer struct {
	ctx scope.Context
}

func NewMockDialer(ctx scope.Context) *MockDialer {
	return &MockDialer{
		ctx: ctx.Fork(),
	}
}

func (d *MockDialer) Dial() (mproto.Connection, error) {
	return &MockConnection{
		ctx:      d.ctx.Fork(),
		incoming: make(chan *proto.Packet),
		outgoing: make(chan *proto.Packet),
	}, nil
}

func (conn *MockConnection) Close() error {
	close(conn.outgoing)
	close(conn.incoming)
	conn.ctx.Cancel()
	return nil
}

func (conn *MockConnection) Receive() (*proto.Packet, error) {
	return <-conn.incoming, nil
}

func (conn *MockConnection) Send(p *proto.Packet) error {
	conn.outgoing <- p
	return nil
}
