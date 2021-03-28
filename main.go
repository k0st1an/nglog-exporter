package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var logsChannel chan []byte

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	flag.BoolVar(&conf.PrintRawNginxLogs, "r", false, "print RAW nginx logs")
	flag.BoolVar(&conf.PrintNginxLogs, "v", false, "print nginx logs")
	flag.BoolVar(&conf.PrintErrors, "e", false, "print errors")
	flag.StringVar(&conf.UDPSrv.Addr, "u", "127.0.0.1:8888", "bind to address of UDP server")
	flag.IntVar(&conf.UDPSrv.ReadBuf, "b", 524288, "size the buffer incoming data of UDP server")
	flag.StringVar(&conf.WebMetricsAddr, "m", ":9999", "bind to address of metrics server")
	flag.IntVar(&conf.Parse.Workers, "w", 5, "workers for process logs")
	flag.IntVar(&conf.QueueSize, "q", 4096, "queue size of input logs")
	flag.Parse()

	log.Println("Print RAW nginx logs:", conf.PrintRawNginxLogs)
	log.Println("Print nginx logs:", conf.PrintNginxLogs)
	log.Println("Print error:", conf.PrintErrors)
	log.Println("Queue size:", conf.QueueSize)
	log.Println("Workers for parsing logs:", conf.Parse.Workers)
	log.Println("Buffer size:", conf.UDPSrv.ReadBuf)

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

	logsChannel = make(chan []byte, conf.QueueSize)

	for i := 0; i < conf.Parse.Workers; i++ {
		go logsProcess()
	}

	go internalMetricsProcess()
	go udpServer()

	log.Print("Metrics endpoint: /metrics")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(conf.WebMetricsAddr, nil))
}
