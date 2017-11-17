package handler

import (
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/api"
)

type BaseMessageHandler struct {
	BaseMessageHandlerWrap
}

func (handler *BaseMessageHandler) Handle(packet protocol.Packet, conn api.Conn) {
	msg := handler.Decode(packet, conn)
	if msg != nil {
		msg.DecodeBody()

		//var handshakeMsg = msg.(message.HandshakeMessage)
		//log.Debug("handshake os name "+handshakeMsg.OsName)
		handler.HandleMessage(msg)
	}
}

type BaseMessageHandlerWrap interface {
	Decode(packet protocol.Packet, connection api.Conn) api.Message

	HandleMessage(msg api.Message)
}
