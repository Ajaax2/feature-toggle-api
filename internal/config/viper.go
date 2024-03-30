package config

import (
	"github.com/spf13/viper"
	"log"
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
