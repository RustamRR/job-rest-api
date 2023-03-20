package postgrestore

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"testing"
)

func TestConfig(t *testing.T) {
	viper.SetConfigFile("./../../../configs/test.toml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		t.Fatal(err)
	}
}

func TestDB(t *testing.T) *gorm.DB {
	config := viper.GetViper()
	db, err := gorm.Open(postgres.Open(config.GetString("dsn")), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
		return nil
	}

	return db
}
