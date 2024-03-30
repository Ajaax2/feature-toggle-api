package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllFeatureToggleHandler
// @Summary Sumário do Endpoint
// @Description Descrição detalhada do endpoint.
// @Tags exemplo
// @Accept  jsonW
// @Produce  json
// @Success 200 {object}  SomeResponse
// @Router /api/v1/toggle [get]
func GetAllFeatureToggleHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET",
	})
}
