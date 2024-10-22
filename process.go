package main

import (
	"fmt"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type nginxLog struct {
	Scheme, Status       string
	RequestTime          string `json:"request_time"`
	UpstreamStatus       string `json:"upstream_status"`
	UpstreamConnectTime  string `json:"upstream_connect_time"`
	UpstreamResponseTime string `json:"upstream_response_time"`
}

func logsProcess() {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var nl nginxLog

	for raw := range logs {
		// `<190>Apr 19 18:52:39 : {"scheme":...` > `{"scheme":...`
		jsonLog := raw[23:]
		if conf.PrintNginxLogs {
			fmt.Println("RAW log:", string(jsonLog))
		}

		if err := json.Unmarshal(jsonLog, &nl); err != nil {
			if conf.PrintErrors {
				fmt.Println("RAW log:", string(raw))
				fmt.Println("JSON log:", string(jsonLog))
				fmt.Println("Error:", err.Error())
			}
			errorsParseTotal.Inc()
			continue
		}

		requestsTotal.WithLabelValues(nl.Scheme).Inc()
		statusTotal.WithLabelValues(nl.Status).Inc()
		upstreamStatusTotal.WithLabelValues(nl.UpstreamStatus).Inc()

		if n, ok := strToFloat64(nl.RequestTime); ok {
			requestTimeHist.WithLabelValues().Observe(n)
		}

		if n, ok := strToFloat64(nl.UpstreamConnectTime); ok {
			upstreamConnectTimeHist.WithLabelValues().Observe(n)
		}

		if n, ok := strToFloat64(nl.UpstreamResponseTime); ok {
			upstreamResposeTimeHist.WithLabelValues().Observe(n)
		}
	}
}

func strToFloat64(str string) (float64, bool) {
	val, err := strconv.ParseFloat(str, 64)

	if err != nil {
		return 0, false
	}

	return val, true
}

func internalMetricsProcess() {
	for {
		ngLogQueue.Set(float64(len(logs)))
		time.Sleep(time.Second * 1)
	}
}
