package router

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Initialize() {

	router := gin.Default()
	InitializeRoutes(router)
	err := router.Run("localhost:" + viper.GetString("server.port"))
	if err != nil {
		return
	}
}
