package api

import (
	"net"
	"time"
	"github.com/zyl0501/go-push-client/push/api/protocol"
)

const (
	STATUS_NEW          byte = 0
	STATUS_CONNECTED    byte = 1
	STATUS_DISCONNECTED byte = 2
)

type Conn interface {
	Init(conn net.Conn)
	GetId() string
	IsConnected() bool
	IsReadTimeout() bool
	IsWriteTimeout() bool
	UpdateLastReadTime()
	UpdateLastWriteTime()
	Close() error
	GetConn() net.Conn
	GetSessionContext() *SessionContext
	SetSessionContext(context SessionContext)
	Send(packet protocol.Packet)
}

type ConnectionManager interface {
	Init()
	Add(connection Conn)
	Get(id string) Conn
	RemoveAndClose(id string) Conn
	GetConnNum() int
	Destroy()
}

type SessionContext struct {
	UserId  string
	Tags      string
	Heartbeat time.Duration
	Cipher0   Cipher
}

type Cipher interface {
	Decrypt(data []byte) ([]byte, error)
	Encrypt(data []byte) ([]byte, error)
}
