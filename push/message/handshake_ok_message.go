package message

import (
	"io"
	"github.com/zyl0501/go-push/common/message"
)

type HandshakeOKMessage struct {
	*message.ByteBufMessage

	serverKey []byte
	heartbeat int32
	sessionId string
	expireTime int64
}

func (msg *HandshakeOKMessage) decodeByteBufMessage(reader io.Reader) {
	msg.serverKey = message.DecodeBytes(reader)
	msg.heartbeat = message.DecodeInt32(reader)
	msg.sessionId = message.DecodeString(reader)
	msg.expireTime = message.DecodeInt64(reader)
}

func (msg *HandshakeOKMessage) encodeByteBufMessage(writer io.Writer) {
	message.EncodeBytes(writer, msg.serverKey)
	message.EncodeInt32(writer, msg.heartbeat)
	message.EncodeString(writer, msg.sessionId)
	message.EncodeInt64(writer, msg.expireTime)
}
