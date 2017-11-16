package push

import (
	"github.com/zyl0501/go-push/api/protocol"
	log "github.com/alecthomas/log4go"
	"io"
	"github.com/zyl0501/go-push/common"
	"github.com/zyl0501/go-push/api"
	"github.com/zyl0501/go-push-client/push/handler"
)

type PushClient struct {
	connClient        *ConnectClient
	messageDispatcher common.MessageDispatcher
}

func (client *PushClient) Init() {
	client.messageDispatcher = common.NewMessageDispatcher()
	client.messageDispatcher.Register(protocol.HANDSHAKE, handler.NewHandshakeOkHandler())
	client.messageDispatcher.Register(protocol.PUSH, handler.NewPushHandler())
}

func (client *PushClient) Start() {
	if client.connClient == nil {
		client.connClient = &ConnectClient{}
		client.connClient.Connect("localhost", 9933)
	}
	serverConn := client.connClient.conn
	go client.listen(serverConn)
}

func (client *PushClient) listen(serverConn api.Conn) {
	conn := serverConn.GetConn()
	head := make([]byte, protocol.HeadLength)
	headReadLen := 0
loop:
	for {
		n, err := conn.Read(head[headReadLen:protocol.HeadLength])
		if err != nil {
			if err == io.EOF {
				log.Error("%s connect error: %v", conn.RemoteAddr().String(), err)
				break loop
			}
		} else {
			if uint32(headReadLen)+uint32(n) < uint32(protocol.HeadLength) {
				headReadLen += n
			} else {
				headReadLen = 0
				packet, bodyLength := protocol.DecodePacket(head)
				readLen := 0
				body := make([]byte, bodyLength)
			bodyLoop:
				for {
					n, err := conn.Read(body[readLen: bodyLength])
					if err != nil {
						if err == io.EOF {
							log.Error("%s connect error: %v", conn.RemoteAddr().String(), err)
							break loop
						} else {
							break bodyLoop
						}
					} else {
						if uint32(readLen)+uint32(n) < bodyLength {
							readLen += n
						} else {
							packet.Body = body
							client.messageDispatcher.OnReceive(packet, serverConn)
							break
						}
					}
				}
			}
		}
	}
}
