package push

import (
	"github.com/zyl0501/go-push/api/push"
	"github.com/zyl0501/go-push/api/protocol"
	"net"
	"time"
	"fmt"
	"reflect"
	"unsafe"
	log "github.com/alecthomas/log4go"
	"io"
	"github.com/zyl0501/go-push/common"
	"github.com/zyl0501/go-push/api"
)

type PushClient struct {
	ConnClient ConnectClient
}

func (client *PushClient) Init() {
}

func (client *PushClient) Start() {
	connectClient.Connect("localhost", 9933)
	serverConn := client.ConnClient.conn
	conn := serverConn.GetConn()

	//var rc chan []byte
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
							client.handler(&packet, &serverConn)
							break
						}
					}
				}
			}
		}
	}
}

func (client *PushClient) handler(packet *protocol.Packet, conn *api.Conn) {
	cmd := packet.Cmd
	switch cmd {
	case protocol.HANDSHAKE:

	}
}

func (client *PushClient) Send(context push.PushContext) (push.PushResult) {
	conn := client.ConnClient.conn

	packet := protocol.Packet{Cmd: protocol.HANDSHAKE}
	packet.Body = nil
	data := protocol.EncodePacket(packet)

	data2 := make([]byte, len(data)*2)
	copy(data2[0:len(data)], data)
	copy(data2[len(data):], data)
	conn.GetConn().Write(data2[0:18])
	conn.GetConn().Write(data2[18:])
	return push.PushResult{}
}

func sender(conn net.TCPConn) {
	packet := protocol.Packet{Cmd: protocol.HANDSHAKE}
	packet.Body = stringToSliceByte("aaa")
	data := protocol.EncodePacket(packet)
	conn.Write(data[0:len(data)-2])
	time.Sleep(time.Second * 10)
	conn.Write(data[len(data)-2:])
	fmt.Println("send over", len(data), string(data))
}

func sender1(conn net.TCPConn) {
	packet := protocol.Packet{Cmd: protocol.HANDSHAKE}
	packet.Body = stringToSliceByte("aaa")
	data := protocol.EncodePacket(packet)

	data2 := make([]byte, len(data)*2)
	copy(data2[0:len(data)], data)
	copy(data2[len(data):], data)
	conn.Write(data2[0:18])
	time.Sleep(time.Second * 10)
	conn.Write(data2[18:])
	fmt.Println("send over", len(data), string(data))

	time.Sleep(time.Second * 10)
}

func stringToSliceByte(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bh))
}
