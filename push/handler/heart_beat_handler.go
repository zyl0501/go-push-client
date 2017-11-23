package handler

import (
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	log "github.com/alecthomas/log4go"
)

type HeartbeatHandler struct {
}

func (handler *HeartbeatHandler) Handle(packet protocol.Packet, conn api.Conn){
	log.Debug(">>> receive heartbeat pong...")
}