package main

import (
	"log"
	"net"
	"runtime"
)

func udpServer() {
	addr, err := net.ResolveUDPAddr("udp", conf.UDPServerAddr)
	if err != nil {
		log.Fatal(err)
	}

	srv, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer srv.Close()
	exit := make(chan struct{})

	for i := 0; i < runtime.NumCPU(); i++ {
		go workers(srv, exit)
	}

	<-exit
}

//
func workers(srv *net.UDPConn, exit chan struct{}) {
	l := 0
	err := error(nil)
	buf := make([]byte, 512)

	for err == nil {
		l, _, err = srv.ReadFromUDP(buf)

		if err != nil {
			errorsReadFromUDPTotal.Inc()

			if conf.PrintErrors {
				log.Print("UDP:", err.Error())
			}
			continue
		}

		if l > 0 {
			logs <- buf[:l]
		}
	}

	exit <- struct{}{}
}
