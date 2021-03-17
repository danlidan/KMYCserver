package game

import (
	"KMYCserver/msg"
	"encoding/binary"

	"github.com/golang/protobuf/proto"
)

func (m *TCPManager) SendRegisterRsp(data *msg.RegisterRsp) {
	mm, _ := proto.Marshal(data)
	sendMsg := make([]byte, 4)

	binary.BigEndian.PutUint16(sendMsg, uint16(2+len(mm)))
	binary.BigEndian.PutUint16(sendMsg[2:], uint16(msg.ProtoId_RegisterRspId))

	sendMsg = append(sendMsg, mm...)
	m.conn.Write(sendMsg)
}

func (m *TCPManager) SendLoginRsp(data *msg.LoginRsp) {
	mm, _ := proto.Marshal(data)
	sendMsg := make([]byte, 4)

	binary.BigEndian.PutUint16(sendMsg, uint16(2+len(mm)))
	binary.BigEndian.PutUint16(sendMsg[2:], uint16(msg.ProtoId_LoginRspId))

	sendMsg = append(sendMsg, mm...)
	m.conn.Write(sendMsg)
}
