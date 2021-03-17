package game

import (
	"KMYCserver/dao"
	"KMYCserver/msg"

	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

//玩家注册
func (m *TCPManager) RecvRegisterReq(data *msg.RegisterReq) {
	redisconn, err := dao.ConnectRedis()
	defer redisconn.Close()
	if err != nil {
		log.Error("conn redis error")
		return
	}

	//不能为空
	if data.Name == "" || data.Pass == "" {
		m.SendRegisterRsp(&msg.RegisterRsp{Success: false})
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

	//发送回馈消息
	m.SendRegisterRsp(&msg.RegisterRsp{Success: true})
}

//玩家登陆
func (m *TCPManager) RecvLoginReq(data *msg.LoginReq) {
	OnlineUsers.Lock()
	redisconn, _ := dao.ConnectRedis()
	defer OnlineUsers.Unlock()
	defer redisconn.Close()

	//已经在线
	if OnlineUsers.Users[data.Name] {
		log.Info("already logged in ", data.Name)
		m.SendLoginRsp(&msg.LoginRsp{
			Success: false,
		})
		return
	}

	//redis获取密码
	TurePass, err := redis.String(redisconn.Do("HGET", dao.NamePassSet, data.Name))
	if err != nil || TurePass == "" || TurePass != data.Pass {
		//密码错误或未注册
		log.Info("wrong pass or no register ", data.Name)
		m.SendLoginRsp(&msg.LoginRsp{
			Success: false,
		})
		return
	}
	//redis获取积分
	TrueRank, err := redis.Int(redisconn.Do("HGET", dao.NameRankSet, data.Name))
	if err != nil {
		log.Error("get rank error ", data.Name)
		m.SendLoginRsp(&msg.LoginRsp{
			Success: false,
		})
		return
	}

	//新建user结构体，关联到TCPManager
	m.user = &User{
		name: data.Name,
		rank: int32(TrueRank),
	}

	//加入OnlineUsers
	OnlineUsers.Users[data.Name] = true

	//发送消息
	m.SendLoginRsp(&msg.LoginRsp{
		Success: true,
		Name:    data.Name,
		Rank:    int32(TrueRank),
	})
}
