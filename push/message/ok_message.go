package message

import (
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"io"
)

type OKMessage struct {
	*ByteBufMessage

	Cmd byte
	Code byte
	Data string
}

func NewOKMessage(packet protocol.Packet, conn api.Conn) *OKMessage {
	baseMessage := BaseMessage{Pkt:packet, Connection: conn}
	byteMessage := ByteBufMessage{BaseMessage: &baseMessage}
	msg := OKMessage{ByteBufMessage: &byteMessage}
	msg.BaseMessageCodec = &msg
	msg.ByteBufMessageCodec = &msg
	return &msg
}

func (message *OKMessage) DecodeByteBufMessage(reader io.Reader) {
	message.Cmd = DecodeByte(reader)
	message.Code = DecodeByte(reader)
	message.Data = DecodeString(reader)
}

func (message *OKMessage) EncodeByteBufMessage(writer io.Writer) {
	EncodeByte(writer, message.Cmd)
	EncodeByte(writer, message.Code)
	EncodeString(writer, message.Data)
}