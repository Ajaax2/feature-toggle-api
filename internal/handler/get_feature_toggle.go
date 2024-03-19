package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Router 			/feature/id [get]
// @Summary 		Create Feature Toggle
// @Description	 	Save fts in Db.
// @Param	 		fts body request.
// @Accept 			json
// @Produce			json
// @Sucess 			200 {object}
func GetAllFeatureToggleHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "GET",
	})
}
