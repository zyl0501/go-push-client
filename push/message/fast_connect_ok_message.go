package message

import (
	"io"
	"time"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/tools"
)

type FastConnectOKMessage struct {
	*ByteBufMessage

	Heartbeat time.Duration
}

func NewFastConnectOKMessage(sessionId uint32, conn api.Conn) *FastConnectOKMessage {
	pkt := protocol.Packet{Cmd:protocol.FAST_CONNECT, SessionId:sessionId}
	baseMessage := BaseMessage{Pkt: pkt, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := FastConnectOKMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func NewFastConnectOKMessage0(packet protocol.Packet, conn api.Conn) *FastConnectOKMessage {
	baseMessage := BaseMessage{Pkt: packet, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := FastConnectOKMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func (message *FastConnectOKMessage) DecodeByteBufMessage(reader io.Reader) {
	message.Heartbeat = tools.MillisecondToDuration(DecodeInt64(reader))
}

func (message *FastConnectOKMessage) EncodeByteBufMessage(writer io.Writer) {
	EncodeInt64(writer, tools.DurationToMillisecond(message.Heartbeat))
}

func (msg *FastConnectOKMessage) Send() {
	msg.sendRaw()
}
