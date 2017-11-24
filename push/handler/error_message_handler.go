package handler

import (
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/message"
	"github.com/zyl0501/go-push-client/push/api"
	log "github.com/alecthomas/log4go"
)

type ErrorMessageHandler struct {
	*BaseMessageHandler
}

func NewErrorMessageHandler() *ErrorMessageHandler {
	baseHandler := &BaseMessageHandler{}
	h := ErrorMessageHandler{BaseMessageHandler: baseHandler}
	h.BaseMessageHandlerWrap = &h
	return &h
}

func (handler *ErrorMessageHandler) Decode(packet protocol.Packet, conn api.Conn) api.Message {
	return message.NewErrorMessage0(packet.SessionId, conn)
}

func (handler *ErrorMessageHandler) HandleMessage(m api.Message) {
	msg := m.(*message.ErrorMessage)
	switch msg.Cmd{
	case protocol.FAST_CONNECT:
		log.Debug(">>> receive fast connect error, message=%v", msg)
	case protocol.HANDSHAKE:
		log.Debug(">>> receive handshake error, message=%v", msg)
	default:
		log.Debug(">>> receive error, message=%v", msg)
	}
}
