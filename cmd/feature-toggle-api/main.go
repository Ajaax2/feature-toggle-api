package main

import (
	"github.com/Ajaax2/feature-toggle-api/internal/config"
	"github.com/Ajaax2/feature-toggle-api/internal/router"
)

// @title 			Feature Toggle API
// @version 		1.0
// @description 	Api responsável por gerenciar os recursos de Feature Toggle do sicredi mobi.
// @host 			localhost:8080
// @contact.name 	Ademar Junior
// @BasePath 		/api/v1
func main() {

	config.Init()
	router.Initialize()

}
