package game

import (
	"KMYCserver/config"
	"net"

	log "github.com/sirupsen/logrus"
)

func UdpConnect() {
	udpAddr, _ := net.ResolveUDPAddr("udp", config.UDPAddr)
	udpConn, _ := net.ListenUDP("udp", udpAddr)
	//defer udpConn.Close()

	log.Info("server ready to read udp from ", config.UDPAddr)

	UdpMgr.conn = udpConn

	go UdpMgr.Receive()
}
