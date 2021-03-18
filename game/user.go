package game

type User struct {
	connManager *TCPManager
	name        string
	rank        int32
	IsMatching  bool //正在匹配
	player      *Player
}
