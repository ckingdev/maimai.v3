package handlers

import (
	"encoding/json"
	"fmt"

	mproto "github.com/cpalone/maimai.v3/proto"

	"euphoria.io/heim/proto"
	"euphoria.io/scope"
)

//type Handler interface {
//	Run(ctx scope.Context) error
//	HandleIncoming(conn Connection, p *proto.Packet) error
//}

type PingHandler struct{}

func (h *PingHandler) Run(ctx scope.Context) error {
	return nil
}

func (h *PingHandler) HandleIncoming(conn mproto.Connection, p *proto.Packet) error {
	if p.Type != proto.PingEventType {
		return nil
	}
	payload, err := p.Payload()
	if err != nil {
		return err
	}
	pingEvent, ok := payload.(*proto.PingEvent)
	if !ok {
		return fmt.Errorf("Error asserting ping-event as such")
	}
	pingReply := &proto.PingReply{
		UnixTime: pingEvent.UnixTime,
	}
	marshalled, err := json.Marshal(pingReply)
	if err != nil {
		return err
	}
	return conn.Send(&proto.Packet{
		Type: proto.PingReplyType,
		Data: marshalled,
	})
}
