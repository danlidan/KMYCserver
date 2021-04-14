package config

import "time"

const (
	//172.17.16.14
	TCPAddr  = "172.17.16.14:19312"
	UDPAddr  = "172.17.16.14:19313"
	LogicGap = time.Millisecond * 50 //逻辑帧的间隔
)
