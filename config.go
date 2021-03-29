package main

type config struct {
	PrintNginxLogs bool
	PrintErrors    bool
	UDPServerAddr  string
	WebMetricsAddr string
	QueueSize      int
}

var conf config
