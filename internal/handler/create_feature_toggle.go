package handler

import (
	"github.com/gin-gonic/gin"
)

// @Router 			/feature [post]
// @Summary 		Create Feature Toggle
// @Description	 	Save fts in Db.
// @Param	 		fts body request.
// @Accept 			json
// @Produce			json
// @Sucess 			200 {object}
func CreateFeatureToggleHandler(ctx *gin.Context) {
	request := struct {
		role string
	}{}
	ctx.BindJSON(&request)

	logger.Infof("request received: %+v", request)
}
