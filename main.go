package main

import (
	"flag"
	"log"
	"net/http"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var logs chan []byte

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	flag.BoolVar(&conf.PrintNginxLogs, "v", false, "print nginx logs")
	flag.BoolVar(&conf.PrintErrors, "e", false, "print errors")
	flag.StringVar(&conf.UDPSrv.Addr, "u", "127.0.0.1:8888", "bind to address of UDP server")
	flag.StringVar(&conf.WebMetricsAddr, "m", "127.0.0.1:9999", "bind to address of metrics server")
	flag.IntVar(&conf.QueueSize, "q", 4096, "queue size of input logs")
	flag.Parse()

	prometheus.MustRegister(
		parseErrorTotal,
		statusTotal,
		requestsTotal,
		requestTimeHist,
		upstreamStatusTotal,
		upstreamConnectTimeHist,
		upstreamResposeTimeHist,
		ngLogQueue,
	)

	logs = make(chan []byte, conf.QueueSize)

	for i := 0; i < runtime.NumCPU(); i++ {
		go logsProcess()
	}

	go internalMetricsProcess()
	udpServer()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(conf.WebMetricsAddr, nil))
}
