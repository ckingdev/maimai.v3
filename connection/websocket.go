package connection

import (
	"encoding/json"

	mproto "github.com/cpalone/maimai.v3/proto"

	"euphoria.io/heim/proto"
	"euphoria.io/scope"
	"github.com/gorilla/websocket"
)

type WSConnection struct {
	ctx  scope.Context
	conn *websocket.Conn
}

func (conn *WSConnection) Receive() (*proto.Packet, error) {
	var p proto.Packet
	if err := conn.conn.ReadJSON(&p); err != nil {
		return nil, err
	}
	return &p, nil
}

func (conn *WSConnection) Send(p *proto.Packet) error {
	return conn.conn.WriteJSON(p)
}

func (conn *WSConnection) Close() error {
	return conn.conn.Close()
}

type WSDialer struct {
	url string
	ctx scope.Context
}

func (d *WSDialer) Dial() (mproto.Connection, error) {
	conn, _, err := websocket.DefaultDialer.Dial(d.url, nil)
	if err != nil {
		return nil, err
	}
	nickCommand := proto.NickCommand{
		Name: "MaiMai.v3",
	}
	nickPayload, err := json.Marshal(nickCommand)
	if err != nil {
		return nil, err
	}
	if err := conn.WriteJSON(&proto.Packet{
		Type: proto.NickType,
		Data: nickPayload,
	}); err != nil {
		return nil, err
	}
	return &WSConnection{
		ctx:  d.ctx.Fork(),
		conn: conn,
	}, nil
}

func NewWSDialer(ctx scope.Context, url string) *WSDialer {
	return &WSDialer{
		url: url,
		ctx: ctx,
	}
}
