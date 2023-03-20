package main

import (
	"errors"
	"fmt"
	"github.com/RustamRR/job-rest-api/internal/app/apiserver"
	"github.com/spf13/viper"
)

var configInitializedError error = errors.New("не удалось получить конфигурацию приложения")

func main() {
	if err := initConfig(); err != nil {
		return
	}
	apiserver.Start()
}

func initConfig() error {
	viper.SetConfigFile("./configs/server.toml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("error read config: %s", err.Error())
		return configInitializedError
	}

	return nil
}
