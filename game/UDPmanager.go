package game

import (
	"KMYCserver/msg"
	"encoding/binary"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type UDPmanager struct {
	conn *net.UDPConn
}

// len 2Byte | protobuf message
// len 为 message的长度，不包括len本身
// len 为大端编码

//接收的内容始终为nextFrameOpts类型
func (m *UDPmanager) Receive() {
	Allbuffer := make([]byte, 4096)

	for {
		msglen, remoteAddr, err := m.conn.ReadFromUDP(Allbuffer)
		if err != nil {
			log.Error("wrong udp read", err)
		}

		curidx := 0
		for curidx < msglen {
			len := binary.BigEndian.Uint16(Allbuffer[curidx : curidx+2])
			buffer := Allbuffer[curidx : curidx+int(len)+2]
			curidx += int(len) + 2

			resData := &msg.Frame{}
			proto.Unmarshal(buffer[2:], resData)

			go m.testSend(buffer, remoteAddr)

			log.Info("recv udp msg ", resData.TestInfo, remoteAddr)
		}
	}
}

func (m *UDPmanager) testSend(b []byte, addr *net.UDPAddr) {
	time.Sleep(time.Second * 3)
	log.Info("send message back to addr ", addr)
	m.conn.WriteToUDP(b, addr)
}
