package handler

import (
	"github.com/gin-gonic/gin"
)

func CreateFeatureToggleHandler(ctx *gin.Context) {
	request := struct {
		role string
	}{}
	ctx.BindJSON(&request)
}
