package main

import (
	"KMYCserver/game"

	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("KYMC server begin")
	//begin udp listen
	game.UdpConnect()
	//begin tcp listen
	game.TcpConnect()
}
