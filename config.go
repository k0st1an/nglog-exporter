package main

type config struct {
	PrintNginxLogs bool
	PrintErrors    bool
	WebMetricsAddr string
	QueueSize      int
	UDPSrv         struct {
		Addr string
	}
}

var conf config
