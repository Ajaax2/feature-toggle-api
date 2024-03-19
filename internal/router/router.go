package router

import (
	"strconv"

	config "github.com/Ajaax2/feature-toggle-api/configs"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	port := strconv.Itoa(config.GetServerPort())
	router := gin.Default()
	InitializeRoutes(router)
	router.Run(port)
}
