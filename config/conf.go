package config

import "time"

const (
	TCPAddr  = "127.0.0.1:19312"
	UDPAddr  = "127.0.0.1:19313"
	LogicGap = time.Millisecond * 50 //逻辑帧的间隔
)
