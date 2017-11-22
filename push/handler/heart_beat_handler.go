package handler

import (
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	log "github.com/alecthomas/log4go"
	"bufio"
	"time"
)

type HeartbeatHandler struct {
}

func (handler *HeartbeatHandler) Handle(packet protocol.Packet, conn api.Conn){
	log.Debug(">>> receive heartbeat pong...")

	go func() {
		t := time.Tick(conn.GetSessionContext().Heartbeat)
		<-t
		writer := bufio.NewWriter(conn.GetConn())
		writer.Write(protocol.EncodePacket(protocol.Packet{Cmd:protocol.HEARTBEAT}))
		writer.Flush()
	}()
}