package config

// Config ...
type Config struct {
	Debug          bool
	WebMetricsAddr string
	UDPSrv         struct {
		Port    int
		ReadBuf int
	}
	Parse struct {
		CutToFirst string
		LogFormat  string
		Workers    int
	}
}

// Conf ...
var Conf Config
