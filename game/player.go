package game

import (
	"KMYCserver/msg"
	"net"
)

type Player struct {
	user        *User
	playerId    int32
	room        *Room
	syncId      int32         //该玩家同步到的操作
	udpAddr     *net.UDPAddr  //为nil说明掉掉线
	matchInfo   *msg.MatchRsp //用于掉线重连发送的消息
	lastSendNum int32         //上一次发送的帧数
}
