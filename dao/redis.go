package dao

import "github.com/gomodule/redigo/redis"

const (
	addr        = "127.0.0.1:6379"
	LoggedInSet = "KMYCLogIn"
	NameRankSet = "KMYCUserData"
	NamePassSet = "KMYCUserPass"
)

func ConnectRedis() (redis.Conn, error) {
	conn, err := redis.Dial("tcp", addr)
	return conn, err
}
