package config

type applicationConfig struct {
	Port int
}

type httpConfigProperty struct {
	MaxIdleConnectionsPerHost int
	DialerTimeout             int
	ConfigAPITimeoutMillis    int
	CommentAPITimeoutMillis   int
	CommentAPIURL             string
}
