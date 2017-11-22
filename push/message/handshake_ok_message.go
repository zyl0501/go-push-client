package message

import (
	"io"
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"time"
	"github.com/zyl0501/go-push-client/push/tools"
)

type HandshakeOKMessage struct {
	*ByteBufMessage

	ServerKey  []byte
	Heartbeat  time.Duration
	SessionId  string
	ExpireTime int64
}

func NewHandshakeOKMessage(packet protocol.Packet, conn api.Conn) *HandshakeOKMessage {
	baseMessage := BaseMessage{Pkt:packet, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := HandshakeOKMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func NewHandshakeOKMessage0(conn api.Conn) *HandshakeOKMessage {
	packet := protocol.Packet{Cmd:protocol.HANDSHAKE, SessionId:protocol.GetSessionId()}
	baseMessage := BaseMessage{Pkt:packet, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := HandshakeOKMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func (msg *HandshakeOKMessage) DecodeByteBufMessage(reader io.Reader) {
	msg.ServerKey = DecodeBytes(reader)
	msg.Heartbeat = tools.MillisecondToDuration(DecodeInt64(reader))
	msg.SessionId = DecodeString(reader)
	msg.ExpireTime = DecodeInt64(reader)
}

func (msg *HandshakeOKMessage) EncodeByteBufMessage(writer io.Writer) {
	EncodeBytes(writer, msg.ServerKey)
	EncodeInt64(writer, tools.DurationToMillisecond(msg.Heartbeat))
	EncodeString(writer, msg.SessionId)
	EncodeInt64(writer, msg.ExpireTime)
}
