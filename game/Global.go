package game

import "sync"

type UsersMap struct {
	sync.Mutex
	Users map[string]bool
}

type MatchPool struct {
	sync.Mutex
	Users []*User
}

var (
	//保存所有在线角色的map， name->bool
	OnlineUsers = &UsersMap{
		Users: make(map[string]bool),
	}
	//匹配池，目前为简单机制
	MatchUsers = &MatchPool{
		Users: make([]*User, 0),
	}
)
