package game

import (
	"KMYCserver/msg"
	"net"

	log "github.com/sirupsen/logrus"
)

func RecvNextFrameOpts(data *msg.NextFrameOpts, addr *net.UDPAddr) {
	//寻找对应房间并执行
	RoomsMap.RLock()
	room := RoomsMap.Rooms[data.RoomId]
	RoomsMap.RUnlock()

	if room == nil {
		log.Error("on recv next Frame opts get room id error ", data.RoomId)
	}

	room.RecvNextFrameOpts(data, addr)
}
