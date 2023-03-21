package apiserver

import (
	"errors"
	"github.com/RustamRR/job-rest-api/internal/store/postgrestore"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"
)

var cannotStartServer error = errors.New("cannot start server")

func Start() {
	config := viper.GetViper()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	gormLogger := zapgorm2.New(logger)
	gormLogger.SetAsDefault()

	db, err := gorm.Open(
		postgres.Open(
			config.GetString("dsn"),
		),
		&gorm.Config{Logger: gormLogger},
	)

	if err != nil {
		logger.Fatal(cannotStartServer.Error())
		return
	}

	store := postgrestore.New(db)

	server := New(logger, store)
	server.Run(config.GetString("server_port"))
}
