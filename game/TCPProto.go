package game

import (
	"KMYCserver/dao"
	"KMYCserver/msg"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

func (m *TCPManager) RecvRegisterReq(data *msg.RegisterReq) {
	redisconn, err := dao.ConnectRedis()
	if err != nil {
		log.Error("conn redis error")
		return
	}

	exist, err := redis.Int(redisconn.Do("HEXISTS", dao.NameRankSet, data.Name))
	if err != nil || exist != 0 {
		log.Error("already registered ", data.Name)
		m.SendRegisterRsp(&msg.RegisterRsp{Success: false})
		return
	}

	//保存密码和rank分（简易版）
	redisconn.Do("HSET", dao.NamePassSet, data.Name, data.Pass)
	redisconn.Do("HSET", dao.NameRankSet, data.Name, "1200")
	log.Info("successfully register ", data.Name)

	m.SendRegisterRsp(&msg.RegisterRsp{Success: true})
}
