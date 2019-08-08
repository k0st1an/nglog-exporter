package metrics

import "github.com/prometheus/client_golang/prometheus"

// ParseErrorTotal ...
var ParseErrorTotal = prometheus.NewCounter(prometheus.CounterOpts{
	Name: "nglog_parse_errors_total",
})

// HTTPRequestTotal ...
var HTTPRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_http_requests_total",
}, []string{"method"})

// StatusTotal ...
var StatusTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_status_total",
}, []string{"code"})

// RequestTimeHist ...
var RequestTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_request_time_hist",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{})

// UpstreamStatusTotal ...
var UpstreamStatusTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	Name: "nglog_upstream_status_total",
}, []string{"code"})

// UpstreamConnectTimeHist ...
var UpstreamConnectTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_upstream_connect_time_hist",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{})

// UpstreamResposeTimeHist ...
var UpstreamResposeTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_upstream_response_time_hist",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{})

// UpstreamHeaderTimeHist ...
var UpstreamHeaderTimeHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:    "nglog_upstream_header_time_hist",
	Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{})
