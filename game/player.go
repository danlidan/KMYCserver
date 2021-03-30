package game

import "net"

type Player struct {
	user     *User
	playerId int32
	room     *Room
	syncId   int32        //该玩家同步到的操作
	udpAddr  *net.UDPAddr //为nil说明掉掉线
}
