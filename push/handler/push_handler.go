package handler

import (
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/message"
	log "github.com/alecthomas/log4go"
)

type PushHandler struct {
	*BaseMessageHandler
}

func NewPushHandler() *PushHandler {
	baseHandler := &BaseMessageHandler{}
	handler := PushHandler{BaseMessageHandler: baseHandler}
	handler.BaseMessageHandlerWrap = &handler
	return &handler
}

func (handler *PushHandler) Decode(packet protocol.Packet, conn api.Conn) api.Message {
	msg := message.NewPushUpMessage(packet, conn)
	return msg
}

func (handler *PushHandler) HandleMessage(m api.Message) {
	msg := m.(*message.PushMessage)
	log.Debug("receive push " + string(msg.Content))
}
