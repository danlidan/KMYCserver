package game

import "time"

//room作为一个房间
type Room struct {
	roomId  int32
	players []*Player
	frameId int32 //目前游戏已经进行到的帧数
}

//游戏收集全部帧0后才能开始(收集udp地址)
func (r *Room) OnGameBegin() {
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
	//三秒后开始每隔50ms的广播
	time.Sleep(time.Second * 3)
	go r.OnLogicSend()
}

//用于每50ms同步广播全部客户端操作
func (r *Room) OnLogicSend() {

}
