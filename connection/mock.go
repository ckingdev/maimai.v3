package connection

import (
	mproto "github.com/cpalone/maimai.v3/proto"

	"euphoria.io/heim/proto"
	"euphoria.io/scope"
)

// MockConnection is a type for internal testing. It simply creates a struct
// with incoming and outgoing channels for the test harness to get or put packets
// on.
type MockConnection struct {
	ctx      scope.Context
	incoming chan *proto.Packet
	outgoing chan *proto.Packet
}

// MockDialer satisfies proto.Dialer, providing a way to create MockConnections.
type MockDialer struct {
	ctx scope.Context
}

// NewMockDialer creates a new MockDialer.
func NewMockDialer(ctx scope.Context) *MockDialer {
	return &MockDialer{
		ctx: ctx.Fork(),
	}
}

// Dial creates a new MockConnection.
func (d *MockDialer) Dial() (mproto.Connection, error) {
	return &MockConnection{
		ctx:      d.ctx.Fork(),
		incoming: make(chan *proto.Packet),
		outgoing: make(chan *proto.Packet),
	}, nil
}

// Close closes the test harness channels and cancels the scope.Context.
func (conn *MockConnection) Close() error {
	close(conn.outgoing)
	close(conn.incoming)
	conn.ctx.Cancel()
	return nil
}

// Receive blocks until a packet is put on the incoming channel by the test
// harness.
func (conn *MockConnection) Receive() (*proto.Packet, error) {
	return <-conn.incoming, nil
}

// Send puts the given packet on the outgoing test harness channel.
func (conn *MockConnection) Send(p *proto.Packet) error {
	conn.outgoing <- p
	return nil
}
