package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func Init() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()

	if err != nil {
		return fmt.Errorf("Error on loading configs: %v", err)
	}

	return nil
}
