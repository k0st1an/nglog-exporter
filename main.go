package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/k0st1an/nglog-exporter/config"
	"github.com/k0st1an/nglog-exporter/metrics"
	"github.com/k0st1an/nglog-exporter/parselog"
	"github.com/k0st1an/nglog-exporter/udpsrv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func init() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	flag.BoolVar(&config.Conf.Debug, "debug", false, "debug mode")
	flag.IntVar(&config.Conf.UDPSrv.Port, "udp-server-port", 8888, "port of UDP server")
	flag.IntVar(&config.Conf.UDPSrv.ReadBuf, "udp-server-in-buf", 524288, "size the buffer incoming data of UDP server")
	flag.StringVar(&config.Conf.Parse.CutToFirst, "cut-to-first", "_", "to cut to the first character")
	flag.StringVar(&config.Conf.Parse.LogFormat, "log-format", "_$host $request_method $scheme $status $request_time $upstream_status $upstream_connect_time $upstream_response_time $upstream_header_time", "log_format from nginx")
	flag.StringVar(&config.Conf.WebMetricsAddr, "web-metrics-addr", ":9999", "bind to address of metrics server")
	flag.IntVar(&config.Conf.Parse.Workers, "workers", 5, "workers for process log")
	flag.Parse()

	prometheus.MustRegister(
		metrics.ParseErrorTotal,
		metrics.HTTPRequestTotal,
		metrics.StatusTotal,
		metrics.RequestTimeHist,
		metrics.UpstreamStatusTotal,
		metrics.UpstreamConnectTimeHist,
		metrics.UpstreamResposeTimeHist,
		metrics.UpstreamHeaderTimeHist,
	)
}

func main() {
	var parseLog parselog.Parse
	var udpSrv udpsrv.UDPSrv

	parseLog.Channel = make(chan []byte, 4096)
	parseLog.Run()

	udpSrv.Channel = parseLog.Channel
	go udpSrv.Run()
	defer udpSrv.Stop()

	log.Print("Metrics endpoint: /metrics")
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(config.Conf.WebMetricsAddr, nil))
}
