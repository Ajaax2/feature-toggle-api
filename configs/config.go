package config

import "github.com/spf13/viper"

var (
	logger *Logger
)

func LoadConfig() {
	viper.SetConfigName("config") // Nome do arquivo YML
	viper.AddConfigPath(".")      // Caminho para o arquivo YML
	viper.SetConfigType("yaml")   // Tipo de arquivo de configuração

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func GetServerPort() int {
	return viper.GetInt("server.port")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
