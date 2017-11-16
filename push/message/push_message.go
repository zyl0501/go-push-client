package message

import (
	"github.com/zyl0501/go-push/common/message"
	"github.com/zyl0501/go-push/api"
	"github.com/zyl0501/go-push/api/protocol"
)

type PushMessage struct {
	*message.BaseMessage
	Content []byte
}

func NewPushUpMessage(packet protocol.Packet, conn api.Conn) *PushMessage {
	baseMessage := message.BaseMessage{Pkt: packet, Connection: conn}
	msg := PushMessage{BaseMessage: &baseMessage}
	msg.BaseMessageCodec = &msg
	return &msg
}

func NewPushMessage0(conn api.Conn) *PushMessage {
	packet := protocol.Packet{Cmd: protocol.PUSH, SessionId: protocol.GetSessionId()}
	baseMessage := message.BaseMessage{Pkt: packet, Connection: conn}
	msg := PushMessage{BaseMessage: &baseMessage}
	msg.BaseMessageCodec = &msg
	return &msg
}

func (msg *PushMessage) decodeBaseMessage(body []byte) {
	msg.Content = body
}
func (msg *PushMessage) encodeBaseMessage() ([]byte) {
	return msg.Content
}
