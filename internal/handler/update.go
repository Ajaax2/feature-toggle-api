package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateFeatureToggleHandler
// @Summary Sumário do Endpoint
// @Description Descrição detalhada do endpoint.
// @Tags exemplo
// @Accept  jsonW
// @Produce  json
// @Success 200 {object}  SomeResponse
// @Router /api/v1/toggle [get]
func UpdateFeatureToggleHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"msg": "UPDATE",
	})
}
