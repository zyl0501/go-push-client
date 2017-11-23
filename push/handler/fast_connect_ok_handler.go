package handler

import (
	log "github.com/alecthomas/log4go"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/message"
	"github.com/zyl0501/go-push-client/push/api"
)

type FastConnectOKHandler struct {
	*BaseMessageHandler
}

func NewFastConnectOKHandler() *FastConnectOKHandler {
	baseHandler := &BaseMessageHandler{}
	h := FastConnectOKHandler{BaseMessageHandler: baseHandler}
	h.BaseMessageHandlerWrap = &h
	return &h
}

func (handler *FastConnectOKHandler) Decode(packet protocol.Packet, conn api.Conn) api.Message {
	return message.NewFastConnectOKMessage0(packet, conn)
}

func (handler *FastConnectOKHandler) HandleMessage(m api.Message) {
	msg := m.(*message.FastConnectOKMessage)
	log.Debug(">>> fast connect ok, message=%v", msg)
	msg.GetConnection().GetSessionContext().Heartbeat = msg.Heartbeat
}
