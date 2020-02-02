package config

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func Init() {
	config = viper.New()

	config.SetConfigType("yaml")
	config.SetConfigName("env")
	config.AddConfigPath("../config")
	config.AddConfigPath("config/")

	err := config.ReadInConfig()
	if err != nil {
		log.Fatalf("error on parsing configuration file: %s", err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}
