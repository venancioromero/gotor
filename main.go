package main

import (
	"flag"
	"gotor/config"
	"gotor/constants"
	"gotor/loggerFactory"
	"gotor/torProxy"
)

var (
	logLevel   = ""
	configFile = ""
)

func init() {
	flag.StringVar(&logLevel, "loglevel", "", "Logging level")
	flag.StringVar(&configFile, "config", constants.DEFAULT_CONFIG_FILE_PATH, "Absolute path to config file")
	flag.Parse()
}

func main() {

	conf := config.LoadConfig(configFile)

	if logLevel != "" {
		loggerFactory.SetLoggingLevel(logLevel)
	} else {
		loggerFactory.SetLoggingLevel(conf.Logging.LogLevel)
	}
	logger := loggerFactory.NewLogger()

	tProxy, err := torProxy.NewTorProxy()

	if err != nil {
		logger.Fatal(err)
	}

	bodyResponse, err := tProxy.Get("http://3g2upl4pq6kufc4m.onion/")

	if err != nil {
		logger.Fatal(err)
	}

	if bodyResponse != "" {
		logger.Info(bodyResponse)
	}

}
