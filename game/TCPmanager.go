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

	for {
		msglen, err := m.conn.Read(Allbuffer)
		if err != nil {
			//断线
			log.Info("tcp conn break ", m.conn.RemoteAddr())
			m.UserOffLine()
			m.conn.Close()
			return
		}

		curidx := 0
		for curidx < msglen {
			len := binary.BigEndian.Uint16(Allbuffer[curidx : curidx+2])
			buffer := Allbuffer[curidx : curidx+int(len)+2]
			curidx += int(len) + 2

			id := binary.BigEndian.Uint16(buffer[2:4])

			log.Info("receive msg ", buffer)

			//获取协议的id
			switch id {
			case uint16(msg.ProtoId_RegisterReqId):
				resData := &msg.RegisterReq{}
				proto.Unmarshal(buffer[4:], resData)
				log.Info("register req name ", resData.Name, " pass ", resData.Pass)
				go m.RecvRegisterReq(resData)
			case uint16(msg.ProtoId_LoginReqId):
				resData := &msg.LoginReq{}
				proto.Unmarshal(buffer[4:], resData)
				log.Info("login req name ", resData.Name, " pass ", resData.Pass)
				go m.RecvLoginReq(resData)
			case uint16(msg.ProtoId_MatchReqId):
				resData := &msg.MatchReq{}
				proto.Unmarshal(buffer[4:], resData)
				log.Info("match req type ", resData.Type)
				go m.RecvMatchReq(resData)
			case uint16(msg.ProtoId_MatchCancelReqId):
				resData := &msg.MatchCancelReq{}
				proto.Unmarshal(buffer[4:], resData)
				log.Info("match cancel req")
				go m.RecvMatchCancelReq(resData)
			default:
				log.Error("wrong proto id ", id)
			}
		}
	}
}

//处理玩家离线
//to do : 房间的处理方式
func (m *TCPManager) UserOffLine() {
	//若已经登录
	if m.user != nil {
		//总之设为false
		m.user.IsMatching = false

		//正在游戏中
		if m.user.player != nil {
			m.user.player.room.Lock()
			m.user.player.udpAddr = nil
			m.user.player.syncId = -1 //重新连接将帧号变为-1

			//判断是否房间所有人已经掉线
			noOne := true
			for _, p := range m.user.player.room.players {
				if p.udpAddr != nil {
					noOne = false
				}
			}
			if noOne {
				//房间结束游戏
				m.user.player.room.endGame = true
			}
			m.user.player.room.Unlock()

			//从roomsmap中去除该房间
			if noOne {
				log.Info("delete room ", m.user.player.room.roomId)
				RoomsMap.Lock()
				delete(RoomsMap.Rooms, m.user.player.room.roomId)
				RoomsMap.Unlock()
				//OnlineUser中去除该房间所有玩家
				OnlineUsers.Lock()
				for _, p := range m.user.player.room.players {
					log.Info("delete usr ", p.user.name)
					delete(OnlineUsers.Users, p.user.name)
				}
				OnlineUsers.Unlock()
			}
		} else {
			//OnlineUser中去除之
			OnlineUsers.Lock()
			log.Info("delete user ", m.user.name)
			delete(OnlineUsers.Users, m.user.name)
			OnlineUsers.Unlock()
		}
	}
}
