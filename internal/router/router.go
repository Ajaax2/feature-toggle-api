package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	SERVER_PORT = "SERVER_PORT"
)

func Initialize() {
	port := os.Getenv(SERVER_PORT)
	router := gin.Default()
	InitializeRoutes(router)
	router.Run(port)
}

