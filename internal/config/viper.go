package config

import (
	"log"

	"github.com/spf13/viper"
)

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}
