package message

import (
	"io"
	"time"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/tools"
)

type FastConnectMessage struct {
	*ByteBufMessage

	SessionId       string
	DeviceId        string
	ExpireHeartbeat time.Duration
}

func NewFastConnectMessage(packet protocol.Packet, conn api.Conn) *FastConnectMessage {
	//pkt := protocol.Packet{Cmd:protocol.HANDSHAKE, SessionId:packet.SessionId}
	baseMessage := BaseMessage{Pkt: packet, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := FastConnectMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func NewFastConnectMessage0(conn api.Conn) *FastConnectMessage {
	packet := protocol.Packet{Cmd: protocol.FAST_CONNECT, SessionId: protocol.GetSessionId()}
	baseMessage := BaseMessage{Pkt: packet, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := FastConnectMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func (message *FastConnectMessage) DecodeByteBufMessage(reader io.Reader) {
	message.SessionId = DecodeString(reader)
	message.DeviceId = DecodeString(reader)
	message.ExpireHeartbeat = tools.MillisecondToDuration(DecodeInt64(reader))
}

func (message *FastConnectMessage) EncodeByteBufMessage(writer io.Writer) {
	EncodeString(writer, message.SessionId)
	EncodeString(writer, message.DeviceId)
	EncodeInt64(writer, tools.DurationToMillisecond(message.ExpireHeartbeat))
}
