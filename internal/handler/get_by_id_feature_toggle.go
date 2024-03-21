package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @router 			/feature [get]
// @summary 		Create Feature Toggle
// @description	 	Save fts in Db.
// @param	 		fts body request.
// @accept 			json
// @produce			json
// @sucess 			200 {object}
func GetByIdFeatureToggleHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET",
	})

}
