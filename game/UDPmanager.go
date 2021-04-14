package game

import (
	"KMYCserver/msg"
	"encoding/binary"
	"net"

	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

type UDPmanager struct {
	conn *net.UDPConn
}

// len 4Byte | protobuf message
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
			len := binary.BigEndian.Uint32(Allbuffer[curidx : curidx+4])
			buffer := Allbuffer[curidx : curidx+int(len)+4]
			curidx += int(len) + 4

			resData := &msg.NextFrameOpts{}
			proto.Unmarshal(buffer[4:], resData)

			go RecvNextFrameOpts(resData, remoteAddr)
		}
	}
}

/*
func (m *UDPmanager) testSend(b []byte, addr *net.UDPAddr) {
	time.Sleep(time.Second * 3)
	log.Info("send message back to addr ", addr)
	m.conn.WriteToUDP(b, addr)
}
*/
