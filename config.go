package main

type config struct {
	PrintRawNginxLogs bool
	PrintNginxLogs    bool
	PrintErrors       bool
	WebMetricsAddr    string
	QueueSize         int
	UDPSrv            struct {
		Addr    string
		ReadBuf int
	}
	Parse struct {
		Workers int
	}
}

var conf config
