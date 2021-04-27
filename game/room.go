package game

import (
	"KMYCserver/config"
	"KMYCserver/msg"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

//room作为一个房间
type Room struct {
	sync.Mutex
	roomId      int32
	playerNum   int
	players     []*Player
	frameId     int32            //目前游戏已经进行到的帧数
	matchFrames []*msg.FrameOpts //保存游戏目前为止的所有帧，下标对应帧号
	nextFrame   *msg.FrameOpts   //下一帧的所有操作
	endGame     bool
}

//游戏收集全部帧0后才能开始(收集udp地址)
func (r *Room) OnGameBegin() {
	/*
		//每个50ms检查一次全部玩家是否存在udp地址
		for {
			for _, p := range r.players {
				if p.udpAddr == nil {
					time.Sleep(time.Millisecond * 50)
					continue
				}
			}
			break
		}
	*/

	//一秒后开始每隔50ms的广播
	//初始化
	r.frameId = 0
	r.matchFrames = make([]*msg.FrameOpts, 0)
	r.nextFrame = &msg.FrameOpts{FrameId: r.frameId, Opts: make([]*msg.OptionEvent, 0)}

	go func() {
		time.Sleep(time.Second * 5)

		log.Info("game begin ", r.roomId)
		//逻辑帧间隔:50ms
		ticker := time.NewTicker(config.LogicGap)
		for _ = range ticker.C {
			r.Lock()
			if r.endGame {
				r.EndGame()
				r.Unlock()
				break
			}
			r.OnLogicSend()
			r.Unlock()
		}
	}()
}

func (r *Room) EndGame() {
	//设置user的player为nil
	for _, p := range r.players {
		u := p.user
		u.player = nil
	}
	//从roommap中删除这个房间
	log.Info("delete room ", r.roomId)
	RoomsMap.Lock()
	delete(RoomsMap.Rooms, r.roomId)
	RoomsMap.Unlock()
}

//用于每50ms同步广播全部客户端操作
func (r *Room) OnLogicSend() {
	r.matchFrames = append(r.matchFrames, r.nextFrame)

	//log.Info("send frame ", r.frameId)
	//发送未同步的帧
	for _, p := range r.players {
		if p.udpAddr != nil {
			//如果上一帧发送小于100帧则正常发送
			if p.lastSendNum < 50 {
				r.SendUnsyncFrames(p)
			} else {
				log.Info(p.user.name, "send last frame num", p.lastSendNum)
				//否则等待一段时间后发送
				p.lastSendNum -= 50
			}
		}
	}

	//当前帧加一
	r.frameId++
	if r.frameId >= int32(MaxFrameNum+5) {
		//游戏抵达时限，结束
		log.Info("time end room ", r.frameId)
		r.endGame = true
		return
	}
	//初始化新的下一帧
	r.nextFrame = &msg.FrameOpts{FrameId: r.frameId, Opts: make([]*msg.OptionEvent, 0)}
}

//对每个玩家发送未同步的帧
func (r *Room) SendUnsyncFrames(p *Player) {
	optFrames := make([]*msg.FrameOpts, 0)

	//最多发送到第几帧，为了让udp发送顺利
	sendMaxFrame := p.syncId + 500
	if sendMaxFrame >= r.frameId {
		sendMaxFrame = r.frameId
	}

	//从 syncid+1 到 frameid 发送帧
	for i := p.syncId + 1; i <= sendMaxFrame; i++ {
		optFrames = append(optFrames, r.matchFrames[i])
	}

	body := &msg.LogicFrame{
		FrameId:      sendMaxFrame,
		UnsyncFrames: optFrames,
	}

	//修改上一帧发送的帧数
	p.lastSendNum = sendMaxFrame - p.syncId - 1
	UdpMgr.SendLogicFrame(body, p.udpAddr)
}

//接收来自一个玩家的下一帧
func (r *Room) RecvNextFrameOpts(data *msg.NextFrameOpts, addr *net.UDPAddr) {
	//需要加锁
	r.Lock()
	defer r.Unlock()

	if data.PlayerId >= int32(r.playerNum) {
		log.Error("on recv next Frame opts get player id error", data.PlayerId, "player num", r.playerNum)
		return
	}

	p := r.players[data.PlayerId]
	//更新玩家当前的udp地址
	p.udpAddr = addr

	//更新该玩家已经同步到的帧数
	if p.syncId < data.FrameId-1 {
		p.syncId = data.FrameId - 1
	}

	//如果帧号不等于当前帧则丢弃
	if data.FrameId != r.frameId {
		return
	}

	//将所有操作加入nextframe中
	for _, opt := range data.Opts {
		r.nextFrame.Opts = append(r.nextFrame.Opts, opt)
	}
}
