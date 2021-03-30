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
	if OnlineUsers.Users[data.Name] != nil {
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
		name:        data.Name,
		rank:        int32(TrueRank),
		connManager: m,
	}

	//加入OnlineUsers
	OnlineUsers.Users[data.Name] = m.user

	//发送消息
	m.SendLoginRsp(&msg.LoginRsp{
		Success: true,
		Name:    data.Name,
		Rank:    int32(TrueRank),
	})
}

func (m *TCPManager) matchV1(playernum int) {
	MatchUsers.Lock()
	defer MatchUsers.Unlock()
	//准备所有匹配的玩家
	UsersToBegin := make([]*User, 0)
	//加入当前玩家
	UsersToBegin = append(UsersToBegin, m.user)
	for len(MatchUsers.Users) > 0 && len(UsersToBegin) < playernum {
		tmpUser := MatchUsers.Users[0]
		//排除所有已经不再匹配的玩家
		if tmpUser.IsMatching {
			UsersToBegin = append(UsersToBegin, tmpUser)
		}
		MatchUsers.Users = MatchUsers.Users[1:]
	}
	//UsersToBegin,或者人数不足，此时matchusers为空，或者人数足够
	if len(UsersToBegin) == playernum {
		//人数足够，通知被匹配的玩家开始
		for _, user := range UsersToBegin {
			user.IsMatching = false
		}
		//todo:通知开始游戏
		//创建对应的room
		newRoom := &Room{
			roomId:    RoomId.GetNewRoomId(),
			playerNum: playernum,
			players:   make([]*Player, playernum),
			frameId:   0, //目前执行到的逻辑帧id
		}

		playerInfo := make([]*msg.PlayerInfoBegin, 0)
		for idx, user := range UsersToBegin {
			//创建对应的player
			user.player = &Player{
				user:     user,
				playerId: int32(idx),
				room:     newRoom,
				syncId:   -1, //初始化已同步的帧id为-1
			}
			playerInfo = append(playerInfo, &msg.PlayerInfoBegin{
				PlayerId: int32(idx),
				Name:     user.name,
				Rank:     user.rank,
			})
			//加入对应的房间
			newRoom.players[idx] = user.player
		}

		//房间加入全局变量Rooms
		RoomsMap.Lock()
		RoomsMap.Rooms[newRoom.roomId] = newRoom
		RoomsMap.Unlock()

		log.Info("begin game roomid ", newRoom.roomId)

		//房间执行开始游戏逻辑
		newRoom.OnGameBegin()

		//通知开始,发送初始信息
		for idx, user := range UsersToBegin {
			user.connManager.SendMatchRsp(&msg.MatchRsp{
				PlayerNum:  int32(playernum),
				MyPlayerId: int32(idx),
				Players:    playerInfo,
				RoomId:     newRoom.roomId,
			})
		}

	} else {
		//人数不足
		m.user.IsMatching = true
		MatchUsers.Users = UsersToBegin
	}

}

//玩家开始匹配
func (m *TCPManager) RecvMatchReq(data *msg.MatchReq) {
	if m.user.IsMatching {
		log.Error("already matching!", m.user.name)
		return
	}

	switch data.Type {
	case 0:
		//1v1，玩家数2人，加入当前玩家
		m.matchV1(2)
	case 1:
		//2v2, 玩家数4人
		m.matchV1(4)
	default:
		log.Error("wrong match req type ", data.Type)
		return
	}
}

//玩家取消匹配
func (m *TCPManager) RecvMatchCancelReq(data *msg.MatchCancelReq) {
	MatchUsers.Lock()
	defer MatchUsers.Unlock()
	//已经开始游戏
	if m.user.player != nil {
		//获得锁后已经开始
		m.SendMatchCancelRsp(&msg.MatchCancelRsp{Success: false})
	} else {
		//获得锁后未开始
		m.user.IsMatching = false
		m.SendMatchCancelRsp(&msg.MatchCancelRsp{Success: true})
	}
}
