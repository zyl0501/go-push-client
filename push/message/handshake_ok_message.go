package message

import (
	"io"
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	log "github.com/alecthomas/log4go"
)

type HandshakeOKMessage struct {
	*ByteBufMessage

	ServerKey  []byte
	Heartbeat  int32
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
	log.Debug("HandshakeOKMessage decodeByteBufMessage")
	msg.ServerKey = DecodeBytes(reader)
	msg.Heartbeat = DecodeInt32(reader)
	msg.SessionId = DecodeString(reader)
	msg.ExpireTime = DecodeInt64(reader)
}

func (msg *HandshakeOKMessage) EncodeByteBufMessage(writer io.Writer) {
	EncodeBytes(writer, msg.ServerKey)
	EncodeInt32(writer, msg.Heartbeat)
	EncodeString(writer, msg.SessionId)
	EncodeInt64(writer, msg.ExpireTime)
}
