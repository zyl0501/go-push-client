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
	"time"
	"context"
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
	client.messageDispatcher.Register(protocol.OK, handler.NewOKMessageHandler())
	client.messageDispatcher.Register(protocol.HEARTBEAT, &handler.HeartbeatHandler{})
}

func (client *PushClient) Start() {
	client.Connect("localhost", 9933)

	connCtx, cancel := context.WithCancel(context.Background())
	go client.listen()
	go client.heartbeatCheck(connCtx, cancel)
	client.handshake()
}

func (client *PushClient) Stop() {}

func (client *PushClient) listen() {
	conn := client.conn.GetConn()
	for {
		packet, err := ReadPacket(conn)
		if err != nil {
			if err == io.EOF {
				log.Error("%s connect error: %v", conn.RemoteAddr().String(), err)
				break
			} else {
				log.Error("%s read error: %v", conn.RemoteAddr().String(), err)
				break
			}
		}
		client.conn.UpdateLastReadTime()
		client.messageDispatcher.OnReceive(*packet, client.conn)
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

func (client *PushClient) reconnect() {
	client.Close()
	client.Start()
}

func (client *PushClient) handshake() {
	ctx := client.conn.GetSessionContext()
	ctx.Cipher0, _ = security.NewRsaCipher()
	handshakeMsg := message.NewHandshakeMessage0(client.conn)
	handshakeMsg.DeviceId = "1111"
	handshakeMsg.OsName = "Windows"
	handshakeMsg.OsVersion = "10"
	handshakeMsg.ClientVersion = "1.0"
	handshakeMsg.Iv = security.CipherBoxIns.RandomAESIV()
	handshakeMsg.ClientKey = security.CipherBoxIns.RandomAESKey()
	handshakeMsg.MinHeartbeat = 5 * time.Second
	handshakeMsg.MaxHeartbeat = 10 * time.Second
	handshakeMsg.Timestamp = 0
	handshakeMsg.Send()
	ctx.Cipher0 = &security.AesCipher{Key: handshakeMsg.ClientKey, Iv: handshakeMsg.Iv}
}

func (client *PushClient) heartbeatCheck(ctx context.Context, cancel context.CancelFunc) {
	conn := client.conn
	hbTimeoutTimes := 0
	for {
		select {
		case t := <-time.After(conn.GetSessionContext().Heartbeat):
			log.Debug("time after result: %v", t)
			if conn.IsReadTimeout() {
				hbTimeoutTimes++
				log.Warn("heartbeat timeout times=%d", hbTimeoutTimes)
			} else {
				hbTimeoutTimes = 0
				log.Debug("connection health for read")
			}
			if hbTimeoutTimes >= MAX_HB_TIMEOUT_COUNT {
				log.Warn("heartbeat timeout times=%d over limit=%d, client restart", hbTimeoutTimes, MAX_HB_TIMEOUT_COUNT)
				hbTimeoutTimes = 0
				client.reconnect()
				cancel()
				continue
			}
			if conn.IsWriteTimeout() {
				log.Debug("<<< send heartbeat ping...")
				conn.Send(protocol.Packet{Cmd: protocol.HEARTBEAT})
			} else {
				log.Debug("connection health for write, next loop send heartbeat")
			}
		case <-ctx.Done():
			log.Info("heartbeat check cancel because of context done.")
			return;
		}
	}
}
