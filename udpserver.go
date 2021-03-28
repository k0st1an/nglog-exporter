package main

import (
	"log"
	"net"
)

func udpServer() {
	addr, err := net.ResolveUDPAddr("udp", conf.UDPSrv.Addr)
	if err != nil {
		log.Fatal(err)
	}

	srv, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer srv.Close()

	if err := srv.SetReadBuffer(conf.UDPSrv.ReadBuf); err != nil {
		log.Fatal(err)
	}

	log.Printf("UDP server started on address %s\n", conf.UDPSrv.Addr)
	buf := make([]byte, 512)

	for {
		rlen, _, err := srv.ReadFromUDP(buf)

		if err != nil {
			log.Print(err)
			continue
		}

		if rlen > 0 {
			// <190>Apr 19 18:52:39 : {"host":... > {"host":...
			// logsChannel <- buf[23:rlen]
			logsChannel <- buf[:rlen]
		}
	}
}
