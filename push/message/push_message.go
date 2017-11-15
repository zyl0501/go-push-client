package message

import "github.com/zyl0501/go-push/common/message"

type PushMessage struct {
	*message.BaseMessage
	content []byte
}

func (msg *PushMessage)decodeBaseMessage(body []byte){
	msg.content = body
}
func (msg *PushMessage) encodeBaseMessage() ([]byte){
	return msg.content
}