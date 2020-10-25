package loggerFactory

import (
	"init_program/constants"
	log "github.com/sirupsen/logrus"
	"strings"
)

var LOG_LEVEL = log.InfoLevel

func SetLoggingLevel(logLevel string) {

	switch strings.ToUpper(logLevel) {
	case constants.TRACE:
		LOG_LEVEL = log.TraceLevel
	case constants.DEBUG:
		LOG_LEVEL = log.DebugLevel
	case constants.INFO:
		LOG_LEVEL = log.InfoLevel
	case constants.WARN:
		LOG_LEVEL = log.WarnLevel
	case constants.ERROR:
		LOG_LEVEL = log.ErrorLevel
	default:
		logger := NewLogger()
		logger.Info(logLevel + "isn't a valid log level. INFO log level set.")
	}
}
func NewLogger() *log.Logger {
	logger := log.New()
	logger.SetLevel(LOG_LEVEL)
	logger.SetFormatter(&log.TextFormatter{FullTimestamp: true})

	return logger
}
