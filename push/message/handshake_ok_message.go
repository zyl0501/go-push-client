package message

import (
	"io"
	"github.com/zyl0501/go-push/common/message"
	"github.com/zyl0501/go-push/api"
	"github.com/zyl0501/go-push/api/protocol"
	log "github.com/alecthomas/log4go"
)

type HandshakeOKMessage struct {
	*message.ByteBufMessage

	ServerKey  []byte
	Heartbeat  int32
	SessionId  string
	ExpireTime int64
}

func NewHandshakeOKMessage(packet protocol.Packet, conn api.Conn) *HandshakeOKMessage {
	baseMessage := message.BaseMessage{Pkt:packet, Connection: conn}
	byteMessage := message.ByteBufMessage{BaseMessage: &baseMessage}
	msg := HandshakeOKMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func NewHandshakeOKMessage0(conn api.Conn) *HandshakeOKMessage {
	packet := protocol.Packet{Cmd:protocol.HANDSHAKE, SessionId:protocol.GetSessionId()}
	baseMessage := message.BaseMessage{Pkt:packet, Connection: conn}
	byteMessage := message.ByteBufMessage{BaseMessage: &baseMessage}
	msg := HandshakeOKMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func (msg *HandshakeOKMessage) DecodeByteBufMessage(reader io.Reader) {
	log.Debug("HandshakeOKMessage decodeByteBufMessage")
	msg.ServerKey = message.DecodeBytes(reader)
	msg.Heartbeat = message.DecodeInt32(reader)
	msg.SessionId = message.DecodeString(reader)
	msg.ExpireTime = message.DecodeInt64(reader)
}

func (msg *HandshakeOKMessage) EncodeByteBufMessage(writer io.Writer) {
	message.EncodeBytes(writer, msg.ServerKey)
	message.EncodeInt32(writer, msg.Heartbeat)
	message.EncodeString(writer, msg.SessionId)
	message.EncodeInt64(writer, msg.ExpireTime)
}
