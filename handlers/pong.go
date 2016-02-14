package handlers

import (
	"encoding/json"
	"fmt"

	mproto "github.com/cpalone/maimai.v3/proto"

	"euphoria.io/heim/proto"
	"euphoria.io/scope"
)

type PongHandler struct{}

func (h *PongHandler) Run(ctx scope.Context) error {
	return nil
}

func (h *PongHandler) HandleIncoming(conn mproto.Connection, p *proto.Packet) error {
	if p.Type != proto.SendEventType {
		return nil
	}
	payload, err := p.Payload()
	if err != nil {
		return err
	}
	sendEvent, ok := payload.(*proto.SendEvent)
	if !ok {
		return fmt.Errorf("Error asserting send-event as such.")
	}
	if sendEvent.Content != "!ping" && sendEvent.Content != "!ping @MaiMai" {
		return nil
	}
	reply := &proto.SendCommand{
		Content: "pong!",
		Parent:  sendEvent.ID,
	}
	marshal, err := json.Marshal(reply)
	if err != nil {
		return err
	}
	return conn.Send(&proto.Packet{
		Type: proto.SendType,
		Data: marshal,
	})
}
