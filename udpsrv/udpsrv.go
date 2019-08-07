package udpsrv

import (
	"log"
	"net"

	"github.com/k0st1an/nglog-exporter/config"
)

// UDPSrv ...
type UDPSrv struct {
	Channel chan []byte
	Conn    *net.UDPConn
}

// Run ...
func (ngLog *UDPSrv) Run() {
	var err error

	ngLog.Conn, err = net.ListenUDP("udp", &net.UDPAddr{
		Port: config.Conf.UDPSrv.Port,
		IP:   net.ParseIP("127.0.0.1"),
	})
	if err != nil {
		log.Fatal(err)
	}

	err = ngLog.Conn.SetReadBuffer(config.Conf.UDPSrv.ReadBuf)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("UDP server started on address %s\n", ngLog.Conn.LocalAddr().String())

	go ngLog.read()
}

func (ngLog *UDPSrv) read() {
	for {
		msgBuf := make([]byte, 124)
		rlen, _, err := ngLog.Conn.ReadFromUDP(msgBuf)

		if err != nil {
			log.Print(err)
			continue
		}

		if rlen > 0 {
			ngLog.Channel <- msgBuf[:rlen]
		}
	}
}

// Stop ...
func (ngLog *UDPSrv) Stop() {
	ngLog.Conn.Close()
	log.Print("UDP server stopped")
}
