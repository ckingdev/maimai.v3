package handlers

import (
	"encoding/json"
	"fmt"

	mproto "github.com/cpalone/maimai.v3/proto"

	"euphoria.io/heim/proto"
	"euphoria.io/scope"
)

// PingHandler responds to incoming ping-events with the proper ping-reply.
type PingHandler struct{}

// Run is a no-op, the PingHandler does not need to run in the background.
func (h *PingHandler) Run(ctx scope.Context) error {
	return nil
}

// HandleIncoming responds to the given ping-event over the given Connection.
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
