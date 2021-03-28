package main

import "github.com/prometheus/client_golang/prometheus"

var errorsReadFromUDPTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "nglog_errors_read_from_udp_total",
})

var errorsParseTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "nglog_errors_parse_total",
})

var requestsTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_requests_total",
	Help: "request method, usually GET or POST and scheme",
}, []string{"scheme"})

var statusTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_status_total",
	Help: "Response status",
}, []string{"code"})

var requestTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_request_time_hist",
	Help:    "Request processing time in seconds with a milliseconds resolution; time elapsed between the first bytes were read from the client and the log write after the last bytes were sent to the client",
	Buckets: []float64{.05, .1, .2, .3, 1, 2},
}, []string{})

var upstreamStatusTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_upstream_status_total",
	Help: "Keeps status code of the response obtained from the upstream server",
}, []string{"code"})

var upstreamConnectTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_upstream_connect_time_hist",
	Help:    "Keeps time spent on establishing a connection with the upstream server",
	Buckets: []float64{.05, .1, .2, .3, 1, 2},
}, []string{})

var upstreamResposeTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_upstream_response_time_hist",
	Help:    "Keeps time spent on receiving the response from the upstream server",
	Buckets: []float64{.05, .1, .2, .3, 1, 2},
}, []string{})

var ngLogQueue = prometheus.NewGauge(prometheus.GaugeOpts{Name: "nglog_queue"})
