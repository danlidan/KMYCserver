package game

import (
	"KMYCserver/msg"
	"encoding/binary"
	"net"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

//TCPManager : 处理tcp收发
type TCPManager struct {
	conn *net.TCPConn
	user *User
}

//protoc --go_out=. xxx.proto
//proto消息格式： len 2Byte | id 2Byte | protobuf message
//len 为 id + message的长度，不包括len本身
//len & id 均为大端编码， 根据id区分消息类型

func (m *TCPManager) Receive() {
	Allbuffer := make([]byte, 4096)

	defer log.Info("tcp conn break ", m.conn.RemoteAddr())

	for {
		msglen, err := m.conn.Read(Allbuffer)
		if err != nil {
			//断线
			return
		}

		curidx := 0
		for curidx < msglen {
			len := binary.BigEndian.Uint16(Allbuffer[curidx : curidx+2])
			buffer := Allbuffer[curidx : curidx+int(len)+2]
			curidx += int(len) + 2

			id := binary.BigEndian.Uint16(buffer[2:4])

			log.Info("receive msg ", buffer)

			switch id {
			case uint16(msg.ProtoId_RegisterReqId):
				resData := &msg.RegisterReq{}
				proto.Unmarshal(buffer[4:], resData)
				log.Info("register req name ", resData.Name, " pass ", resData.Pass)
				go m.RecvRegisterReq(resData)
			}
		}
	}
}

func (m *TCPManager) SendProto() {

}
