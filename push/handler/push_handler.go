package handler

import (
	"github.com/zyl0501/go-push/core/handler"
	"github.com/zyl0501/go-push/api/protocol"
	"github.com/zyl0501/go-push/api"
)

type PushHandler struct {
	*handler.BaseMessageHandler
}

func (handler *PushHandler) Decode(packet protocol.Packet, conn api.Conn) api.Message {
}

func (handler *PushHandler) HandleMessage(msg api.Message){

}