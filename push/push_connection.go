package push

import (
	"net"
	"github.com/zyl0501/go-push-client/push/api"
	"time"
	"github.com/zyl0501/go-push-client/push/security"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"bufio"
)

var (
	lId = 0
)

type PushConnection struct {
	conn          net.Conn
	status        byte
	lastReadTime  time.Time
	lastWriteTime time.Time
	id            string
	context       api.SessionContext
}

func NewPushConnection() (conn *PushConnection) {
	lId++
	conn = &PushConnection{
		conn:          nil,
		status:        api.STATUS_NEW,
		lastReadTime:  time.Unix(0, 0),
		lastWriteTime: time.Unix(0, 0),
		id:            string(lId)}
	return conn
}

func (serverConn *PushConnection) Init(conn net.Conn) {
	serverConn.conn = conn
	serverConn.status = api.STATUS_CONNECTED
	serverConn.lastReadTime = time.Now()
	serverConn.lastWriteTime = time.Now()
	serverConn.context = api.SessionContext{Heartbeat: DEFAULT_HEARTBEAT}
	cipher, _ := security.NewRsaCipher()
	serverConn.context.Cipher0 = cipher
}

func (serverConn *PushConnection) GetId() string {
	return serverConn.id
}

func (serverConn *PushConnection) IsConnected() bool {
	return serverConn.status == api.STATUS_CONNECTED
}

func (serverConn *PushConnection) IsReadTimeout() bool {
	//加超时时间是防止处理速度过快，导致一直超时。
	//timer    		00:00.0		00.05.5	next timer, the 0.5 second means between twice ticker process time
	//lastWriteTime	00:00.1		00.05.1	timeout
	//lastReadTime	00.00.2		00.05.2	timeout
	return time.Since(serverConn.lastReadTime) > (serverConn.context.Heartbeat + time.Second/2)
}

func (serverConn *PushConnection) IsWriteTimeout() bool {
	return time.Since(serverConn.lastWriteTime) > (serverConn.context.Heartbeat - time.Second/2)
}

func (serverConn *PushConnection) UpdateLastReadTime() {
	serverConn.lastReadTime = time.Now()
}
func (serverConn *PushConnection) UpdateLastWriteTime() {
	serverConn.lastWriteTime = time.Now()
}

func (serverConn *PushConnection) Close() error {
	serverConn.status = api.STATUS_DISCONNECTED
	if serverConn.conn != nil {
		return serverConn.conn.Close()
	}
	return nil
}

func (serverConn *PushConnection) GetConn() net.Conn {
	return serverConn.conn
}

func (serverConn *PushConnection) GetSessionContext() *api.SessionContext {
	return &serverConn.context
}

func (serverConn *PushConnection) SetSessionContext(context api.SessionContext) {
	serverConn.context = context
}

func (serverConn *PushConnection) Send(packet protocol.Packet) {
	writer := bufio.NewWriter(serverConn.conn)
	writer.Write(protocol.EncodePacket(packet))
	writer.Flush()
	serverConn.UpdateLastWriteTime()
}

func (serverConn *PushConnection) Reconnect() {
	//serverConn.Close()
}
