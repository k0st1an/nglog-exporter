package parselog

import (
	"log"
	"strconv"
	"strings"

	"github.com/k0st1an/nglog-exporter/config"
	"github.com/k0st1an/nglog-exporter/metrics"

	"github.com/satyrius/gonx"
)

// Parse ...
type Parse struct {
	Channel chan []byte
}

// Run ...
func (parse *Parse) Run() {
	for i := 0; i < config.Conf.Parse.Workers; i++ {
		go parse.process()
	}
}

// $request_length $body_bytes_sent $bytes_sent $upstream_connect_time $upstream_response_time $upstream_header_time $upstream_response_length
func (parse *Parse) process() {
	for {
		if msg, ok := <-parse.Channel; ok {
			data := strings.TrimSpace(string(msg))
			idxCutToFirst := strings.Index(data, config.Conf.Parse.CutToFirst)

			if config.Conf.Debug {
				log.Print("Data: ", data)
				log.Print("Len: ", len(data))
				log.Print("Index: ", idxCutToFirst)
			}

			if idxCutToFirst == -1 {
				metrics.ParseErrorTotal.Inc()
				continue
			}

			metrics.HTTPRequestTotal.WithLabelValues("http").Inc()

			reader := gonx.NewParser(config.Conf.Parse.LogFormat)
			entry, err := reader.ParseString(data[idxCutToFirst:])

			if err != nil {
				metrics.ParseErrorTotal.Inc()
				continue
			}

			if val, err := entry.Field("status"); err == nil {
				metrics.StatusTotal.WithLabelValues(val).Inc()
			}

			if strVal, err := entry.Field("request_time"); err == nil {
				if val, ok := strToFloat64(strVal); ok {
					metrics.RequestTimeHist.WithLabelValues().Observe(val)
				}
			}

			if val, err := entry.Field("upstream_status"); err == nil {
				metrics.UpstreamStatusTotal.WithLabelValues(val).Inc()
			}

			if strVal, err := entry.Field("upstream_connect_time"); err == nil {
				if val, ok := strToFloat64(strVal); ok {
					metrics.UpstreamConnectTimeHist.WithLabelValues().Observe(val)
				}
			}
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
