package apiserver

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Start() {
	config := viper.GetViper()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	server := New(logger)
	server.Run(config.GetString("server_port"))
}
