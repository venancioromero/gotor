package main

import (
	"flag"
	"init_program/config"
	"init_program/constants"
	"init_program/loggerFactory"
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

	if logLevel == "" {
		loggerFactory.SetLoggingLevel(logLevel)
	} else {
		loggerFactory.SetLoggingLevel(conf.Logging.LogLevel)
	}
}
