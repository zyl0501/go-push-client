package handler

import (
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"github.com/zyl0501/go-push-client/push/api"
	log "github.com/alecthomas/log4go"
	message2 "github.com/zyl0501/go-push-client/push/message"
	"github.com/zyl0501/go-push-client/push/security"
)

type HandshakeOkHandler struct {
	*BaseMessageHandler
	sessionStorage api.SessionStorage
}

func NewHandshakeOkHandler(sessionStorage api.SessionStorage) *HandshakeOkHandler {
	baseHandler := &BaseMessageHandler{}
	handler := HandshakeOkHandler{BaseMessageHandler: baseHandler, sessionStorage: sessionStorage}
	handler.BaseMessageHandlerWrap = &handler
	return &handler
}

func (handler *HandshakeOkHandler) Decode(packet protocol.Packet, conn api.Conn) api.Message {
	return message2.NewHandshakeOKMessage(packet, conn)
}

func (handler *HandshakeOkHandler) HandleMessage(m api.Message) {
	msg := m.(*message2.HandshakeOKMessage)
	log.Debug(">>> handshake ok message=%s ", msg)

	conn := msg.Connection
	ctx := conn.GetSessionContext()
	serverKey := msg.ServerKey

	if len(serverKey) != security.CipherBoxIns.AesKeyLength {
		log.Warn("handshake error serverKey invalid message=%s ", msg);
		conn.Close()
		return;
	}
	//设置心跳
	ctx.Heartbeat = msg.Heartbeat

	//更换密钥
	var cp security.AesCipher
	cp = *ctx.Cipher0.(*security.AesCipher)
	sessionKey := security.CipherBoxIns.MixKey(cp.Key, serverKey);
	ctx.Cipher0 = &security.AesCipher{Key: sessionKey, Iv: cp.Iv}

	//触发握手成功事件
	conn.Send(protocol.Packet{Cmd: protocol.HEARTBEAT})
	//保存token
	handler.saveToken(*msg, *ctx)
}

func (handler *HandshakeOkHandler) saveToken(msg message2.HandshakeOKMessage, ctx api.SessionContext) {
	if msg.SessionId == "" || handler.sessionStorage == nil {
		return
	}
	session := api.PersistentSession{};
	session.SessionId = msg.SessionId;
	session.ExpireTime = msg.ExpireTime;
	session.Cipher0 = ctx.Cipher0;
	handler.sessionStorage.SaveSession(session);
}
