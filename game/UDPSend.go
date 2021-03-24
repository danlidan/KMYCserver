package game

import (
	"KMYCserver/msg"
	"encoding/binary"
	"net"

	"google.golang.org/protobuf/proto"
)

// len 2Byte | protobuf message
// len 为 message的长度，不包括len本身
// len 为大端编码

func (m *UDPmanager) SendLogicFrame(data *msg.LogicFrame, addr *net.UDPAddr) {
	mm, _ := proto.Marshal(data)
	sendMsg := make([]byte, 2)

	binary.BigEndian.PutUint16(sendMsg, uint16(len(mm)))

	sendMsg = append(sendMsg, mm...)
	m.conn.WriteToUDP(sendMsg, addr)
}
