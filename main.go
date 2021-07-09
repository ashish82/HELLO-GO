package main

import (
	"HELLO-GO/config"
	"HELLO-GO/constant"
	"HELLO-GO/controller"
	httpConfig "HELLO-GO/http_config"
	"HELLO-GO/utility/logger"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	confPath = kingpin.Flag("configPath", "Path for configuration files").
			Default("./resources").Short('r').String()
	env = kingpin.Flag("environment", "Environment to use for running the application").
		Default("dev").Short('e').Enum("dev", "qa", "staging", "prod")
	log = logger.GetLogger()
)

func main() {

	log.WithFields(logrus.Fields{
		"GoVersion": runtime.Version(),
		"GOOS":      runtime.GOOS,
		"GOARCH":    runtime.GOARCH,
		"NumCPU":    runtime.NumCPU(),
		"GOPATH":    os.Getenv("GOPATH"),
		"GOROOT":    runtime.GOROOT(),
	}).Info("Starting up ... wait for it... *HELLO GO*")
	//parse the config and set tha property file data in Appconfig struct
	kingpin.Parse()
	config.Initialize(*confPath, *env)
	//initialize http connector for individual timeout
	httpConfig.InitHTTPClients()
	log.Info("ConfigAPI timeout :", config.HttpConfigProperty.ConfigAPITimeoutMillis)
	log.Info("Dialer timeout :", config.HttpConfigProperty.DialerTimeout)
	log.Info("MaxIdleConnect timeout :", config.HttpConfigProperty.MaxIdleConnectionsPerHost)
	// server setup and start
	initializeServer()
	log.Info("Started HELLO-GO")
}

func initializeServer() {

	logfile, err := os.OpenFile(constant.LogFilePath+"access.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		log.Fatalln("Failed to create request log file:", err)
	}

	router := gin.New()
	router.Use(logger.LoggerWithWriter(io.MultiWriter(logfile)))
	router.Use(gin.Recovery())
	apiGroup := router.Group("/hello-go")
	{
		apiGroup.GET("health", controller.GetHealthStatus)
		apiGroup.GET("allComments", controller.GetAllComments)

	}
	err = router.Run(":" + strconv.Itoa(config.AppConfig.Port))
	if err != nil {
		fmt.Println(err)
	}
}
