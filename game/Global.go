package game

import "sync"

type UsersMap struct {
	sync.Mutex
	Users map[string]*User
}

type MatchPool struct {
	sync.Mutex
	Users []*User
}

type RoomMgr struct {
	sync.RWMutex
	Rooms map[int32]*Room
}

type RoomIdMgr struct {
	sync.Mutex
	NextRoomId int32
}

var (
	//保存所有在线角色的map， name->bool
	OnlineUsers = &UsersMap{
		Users: make(map[string]*User),
	}
	//匹配池，目前为简单机制
	MatchUsers = &MatchPool{
		Users: make([]*User, 0),
	}
	//全局用的udp管理
	UdpMgr = &UDPmanager{}
	//房间号对房间的映射
	RoomsMap = &RoomMgr{
		Rooms: make(map[int32]*Room),
	}
	//管理生成的房间号
	RoomId = &RoomIdMgr{
		NextRoomId: 10000,
	}
)

//获取一个新的房间号
func (m *RoomIdMgr) GetNewRoomId() int32 {
	m.Lock()
	defer m.Unlock()
	m.NextRoomId++
	return m.NextRoomId
}
