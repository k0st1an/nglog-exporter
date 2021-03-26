package metrics

import "github.com/prometheus/client_golang/prometheus"

// ParseErrorTotal ...
var ParseErrorTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "nglog_parse_errors_total",
})

// HTTPRequestTotal ...
var HTTPRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_http_requests_total",
	Help: "request method, usually GET or POST and scheme",
}, []string{"method", "scheme"})

// StatusTotal ...
var StatusTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_status_total",
	Help: "Response status",
}, []string{"code", "host"})

// RequestTimeHist ...
var RequestTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_request_time_hist",
	Help:    "Request processing time in seconds with a milliseconds resolution; time elapsed between the first bytes were read from the client and the log write after the last bytes were sent to the client",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{})

// UpstreamStatusTotal ...
var UpstreamStatusTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_upstream_status_total",
	Help: "Keeps status code of the response obtained from the upstream server",
}, []string{"code", "host"})

// UpstreamConnectTimeHist ...
var UpstreamConnectTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_upstream_connect_time_hist",
	Help:    "Keeps time spent on establishing a connection with the upstream server",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{})

// UpstreamResposeTimeHist ...
var UpstreamResposeTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_upstream_response_time_hist",
	Help:    "Keeps time spent on receiving the response from the upstream server",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{})

// UpstreamHeaderTimeHist ...
var UpstreamHeaderTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_upstream_header_time_hist",
	Help:    "Keeps time spent on receiving the response header from the upstream server",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{})
