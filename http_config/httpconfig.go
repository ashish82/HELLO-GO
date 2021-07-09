package http_config

import (
	"HELLO-GO/config"
	"HELLO-GO/constant"
	"net"
	"net/http"
	"time"
)

var HTTPClientMap map[string]*http.Client

func InitHTTPClients() {
	HTTPClientMap = make(map[string]*http.Client)
	//custom and common dialer and transport
	dialer := &net.Dialer{
		Timeout: time.Duration(config.HttpConfigProperty.DialerTimeout) * time.Millisecond,
	}
	transport := &http.Transport{
		MaxIdleConnsPerHost: config.HttpConfigProperty.MaxIdleConnectionsPerHost,
		DialContext:         dialer.DialContext,
	}

	//create clients with timeout configuration
	HTTPClientMap[constant.APPConfigAPI] = &http.Client{
		Transport: transport,
		Timeout:   time.Duration(config.HttpConfigProperty.ConfigAPITimeoutMillis) * time.Millisecond,
	}

	HTTPClientMap[constant.CommentAPI] = &http.Client{
		Transport: transport,
		Timeout:   time.Duration(config.HttpConfigProperty.CommentAPITimeoutMillis) * time.Millisecond,
	}

}
