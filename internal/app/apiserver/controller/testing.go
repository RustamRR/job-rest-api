package controller

import (
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T) {
	viper.SetConfigFile("./../../../../../configs/test.toml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		t.Fatal(err)
	}
}
