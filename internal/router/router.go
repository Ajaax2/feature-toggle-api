package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func Init() {

	router := gin.Default()
	InitializeRoutes(router)
	err := router.Run("localhost:" + viper.GetString("server.port"))
	if err != nil {
		log.Fatal(err)
	}
}
