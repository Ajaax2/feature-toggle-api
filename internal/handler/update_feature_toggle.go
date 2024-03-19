package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateFeatureToggleHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"msg": "UPDATE",
	})

}
