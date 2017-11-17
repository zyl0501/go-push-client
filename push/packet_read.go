package push

import (
	"github.com/zyl0501/go-push-client/push/tools"
	"github.com/zyl0501/go-push-client/push/api/protocol"
	"net"
)

func ReadPacket(conn net.Conn) (*protocol.Packet, error){
	header, err := tools.ReadData(conn, uint32(protocol.HeadLength))
	if err != nil {
		return nil, err
	}
	packet, bodyLen := protocol.DecodePacket(header)

	if bodyLen > 0 {
		body, err := tools.ReadData(conn, bodyLen)
		if err != nil {
			return nil, err
		}
		packet.Body = body
	}
	return &packet, nil
}
