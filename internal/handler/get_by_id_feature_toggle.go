package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router 			/feature [get]
// @Summary 		Create Feature Toggle
// @Description	 	Save fts in Db.
// @Param	 		fts body request.
// @Accept 			json
// @Produce			json
// @Sucess 			200 {object}
func GetByIdFeatureToggleHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET",
	})

}
