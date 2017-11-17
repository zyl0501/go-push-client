package push

import (
	"github.com/zyl0501/go-push-client/push/api/protocol"
	log "github.com/alecthomas/log4go"
	"io"
	"github.com/zyl0501/go-push-client/push/api"
	"github.com/zyl0501/go-push-client/push/handler"
	"net"
	"strconv"
	"github.com/zyl0501/go-push-client/push/security"
	"github.com/zyl0501/go-push-client/push/message"
)

type PushClient struct {
	conn              api.Conn
	config            ClientConfig
	messageDispatcher message.MessageDispatcher
}

func (client *PushClient) Init() {
	client.messageDispatcher = message.NewMessageDispatcher()
	client.messageDispatcher.Register(protocol.HANDSHAKE, handler.NewHandshakeOkHandler())
	client.messageDispatcher.Register(protocol.PUSH, handler.NewPushHandler())
}

func (client *PushClient) Start() {
	client.Connect("localhost", 9933)
	serverConn := client.conn
	go client.listen(serverConn)
}

func (client *PushClient) listen(serverConn api.Conn) {
	conn := serverConn.GetConn()
	for {
		packet, err := ReadPacket(conn)
		if err != nil {
			if err == io.EOF {
				log.Error("%s connect error: %v", conn.RemoteAddr().String(), err)
				break
			} else {
				continue
			}
		}
		client.messageDispatcher.OnReceive(*packet, serverConn)
	}
}

func (client *PushClient) BindUser(userId string, tags string) {
	if userId == "" {
		log.Warn("bind user is null")
		return
	}
	ctx := client.conn.GetSessionContext()
	if ctx.UserId != "" {
		//已经绑定
		if ctx.UserId == userId {
			if ctx.Tags == tags {
				return
			}
		} else {
			//切换用户，要先解绑老用户
			client.UnbindUser()
		}
	}

	ctx.UserId = userId
	ctx.Tags = tags
	client.config.UserId = userId
	client.config.Tags = tags

	msg := message.NewBindUserMessage0(client.conn)
	msg.UserId = userId
	msg.Tags = tags

	log.Info("<<< do bind user, userId=%s", userId)
	msg.Send()
}

func (client *PushClient) UnbindUser() {

}

func (client *PushClient) Connect(host string, port int) *net.TCPConn {
	server := host + ":" + strconv.Itoa(port)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	if err != nil {
		log.Error("Fatal error:%s", err)
		return nil
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Error("Fatal error:%s", err)
		return nil
	}
	client.conn = NewPushConnection()
	client.conn.Init(conn)
	client.handshake()
	return conn
}

func (client *PushClient) Close() {
	if client.conn != nil {
		err := client.conn.Close()
		if err != nil {
			log.Error("Fatal error:%s", err)
		}
	}
}

func (client *PushClient) handshake() {
	context := client.conn.GetSessionContext()
	context.Cipher0, _ = security.NewRsaCipher()
	handshakeMsg := message.NewHandshakeMessage0(client.conn)
	handshakeMsg.DeviceId = "1111"
	handshakeMsg.OsName = "Windows"
	handshakeMsg.OsVersion = "10"
	handshakeMsg.ClientVersion = "1.0"
	handshakeMsg.Iv = security.CipherBoxIns.RandomAESIV()
	handshakeMsg.ClientKey = security.CipherBoxIns.RandomAESKey()
	handshakeMsg.MinHeartbeat = 10000
	handshakeMsg.MaxHeartbeat = 10000
	handshakeMsg.Timestamp = 0
	handshakeMsg.Send()
	context.Cipher0 = &security.AesCipher{Key: handshakeMsg.ClientKey, Iv: handshakeMsg.Iv}
}
