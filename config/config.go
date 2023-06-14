package config

import (
	"os"

	"github.com/loghole/tron"
	"go.uber.org/zap"

	_ "github.com/lib/pq" // init postgres driver.
)

const (
	_defaultHTTPPort  = 8080
	_defaultAdminPort = 8082
)

func TronOptions() []tron.Option {
	return []tron.Option{
		tron.AddLogCaller(),
		tron.AddLogStacktrace(zap.PanicLevel.String()),
		tron.WithPublicHTTP(_defaultHTTPPort),
		tron.WithAdminHTTP(_defaultAdminPort),
		tron.WithLoggerConfig(zap.NewDevelopmentConfig()),
		tron.WithLoggerLevel(os.Getenv("LOGGER_LEVEL")),
	}
}
