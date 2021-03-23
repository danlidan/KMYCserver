package game

import (
	"KMYCserver/config"
	"net"

	log "github.com/sirupsen/logrus"
)

func TcpConnect() {
	//定义一个tcp断点
	var tcpAddr *net.TCPAddr
	//通过ResolveTCPAddr实例一个具体的tcp断点
	tcpAddr, _ = net.ResolveTCPAddr("tcp", config.TCPAddr)
	//打开一个tcp断点监听
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()

	log.Info("server ready to read tcp from ", config.TCPAddr)

	for {
		tcpConn, err := tcpListener.AcceptTCP()
		if err != nil {
			log.Error("accept tcp conn error ", err)
		}

		log.Info("accpet tcp conn ", tcpConn.RemoteAddr())
		//对每个连接生成tcpMangaer单开协程处理
		tcpMgr := &TCPManager{
			conn: tcpConn,
		}

		go tcpMgr.Receive()
	}
}
