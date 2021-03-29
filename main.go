package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const version = "2.0.0"

var logs chan []byte

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	flag.BoolVar(&conf.PrintNginxLogs, "v", false, "print nginx logs")
	flag.BoolVar(&conf.PrintErrors, "e", false, "print errors")
	flag.BoolVar(&conf.PrintVersion, "V", false, "print version and exit")
	flag.StringVar(&conf.UDPServerAddr, "u", "127.0.0.1:8888", "bind to address of UDP server")
	flag.StringVar(&conf.WebMetricsAddr, "m", "127.0.0.1:9999", "bind to address of metrics server")
	flag.IntVar(&conf.QueueSize, "q", 4096, "queue size of input logs")
	flag.Parse()

	if conf.PrintVersion {
		fmt.Printf("v%s\n", version)
		os.Exit(0)
	}

	prometheus.MustRegister(
		ngLogQueue,
		statusTotal,
		requestsTotal,
		requestTimeHist,
		errorsParseTotal,
		upstreamStatusTotal,
		errorsReadFromUDPTotal,
		upstreamConnectTimeHist,
		upstreamResposeTimeHist,
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
