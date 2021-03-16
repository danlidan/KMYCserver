package game

type User struct {
	connManager *TCPManager
	name        string
	rank        int32
}
