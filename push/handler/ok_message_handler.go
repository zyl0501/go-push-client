package handler

import (
	log "github.com/alecthomas/log4go"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/message"
	"github.com/zyl0501/go-push-client/push/api"
)

type OKMessageHandler struct {
	*BaseMessageHandler
}

func NewOKMessageHandler() *OKMessageHandler {
	baseHandler := &BaseMessageHandler{}
	h := OKMessageHandler{BaseMessageHandler: baseHandler}
	h.BaseMessageHandlerWrap = &h
	return &h
}

func (handler *OKMessageHandler) Decode(packet protocol.Packet, conn api.Conn) api.Message {
	msg := message.NewOKMessage(packet, conn)
	return msg
}

func (handler *OKMessageHandler) HandleMessage(m api.Message) {
	msg := m.(*message.OKMessage)
	switch msg.Cmd{
	case protocol.BIND:
		log.Debug("receive bind ok ack")
	}

}
