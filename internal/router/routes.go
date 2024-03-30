package router

import (
	_ "github.com/Ajaax2/feature-toggle-api/docs"
	"github.com/Ajaax2/feature-toggle-api/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitializeRoutes(router *gin.Engine) {

	basePath := "/api/v1"
	v1 := router.Group(basePath)
	{
		v1.GET("/toggle", handler.GetAllFeatureToggleHandler)
		v1.GET("/toggle/{id}", handler.GetByIdFeatureToggleHandler)
		v1.POST("/toggle", handler.CreateFeatureToggleHandler)
		v1.PUT("/toggle", handler.UpdateFeatureToggleHandler)
		v1.DELETE("/toggle/{id}", handler.DeleteFeatureToggleHandler)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
