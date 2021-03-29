package main

type config struct {
	PrintVersion   bool
	PrintNginxLogs bool
	PrintErrors    bool
	UDPServerAddr  string
	WebMetricsAddr string
	QueueSize      int
}

var conf config
