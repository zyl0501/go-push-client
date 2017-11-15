package handler

import (
	"github.com/zyl0501/go-push/core/handler"
	"github.com/zyl0501/go-push/api/protocol"
	"github.com/zyl0501/go-push/api"
)

type HandlerShakeOkHandler struct {
	*handler.BaseMessageHandler
}

func (handler *HandlerShakeOkHandler) Decode(packet protocol.Packet, conn api.Conn) api.Message {
}

func (handler *HandlerShakeOkHandler) HandleMessage(msg api.Message){

}