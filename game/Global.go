package game

import "sync"

type UsersMap struct {
	sync.Mutex
	Users map[string]bool
}

var (
	//保存所有在线角色的map， name->bool
	OnlineUsers = &UsersMap{
		Users: make(map[string]bool),
	}
)
