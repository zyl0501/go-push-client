package message

import (
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/api"
	"io"
)

type BindUserMessage struct {
	*ByteBufMessage

	UserId string
	Alias  string
	Tags   string
}

func NewBindUserMessage(packet protocol.Packet, conn api.Conn) *BindUserMessage {
	baseMessage := BaseMessage{Pkt: packet, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := BindUserMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func NewBindUserMessage0(conn api.Conn) *BindUserMessage {
	packet := protocol.Packet{Cmd: protocol.BIND}
	baseMessage := BaseMessage{Pkt: packet, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := BindUserMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func (message *BindUserMessage) DecodeByteBufMessage(reader io.Reader) {
	message.UserId = DecodeString(reader)
	message.Alias = DecodeString(reader)
	message.Tags = DecodeString(reader)
}

func (message *BindUserMessage) EncodeByteBufMessage(writer io.Writer) {
	EncodeString(writer, message.UserId)
	EncodeString(writer, message.Alias)
	EncodeString(writer, message.Tags)
}