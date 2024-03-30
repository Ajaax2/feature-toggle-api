package toggle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// DeleteFeatureToggleHandler
// @Router 			/feature/id [delete]
// @Summary 		Create Feature Toggle
// @Description	 	Save fts in Db.
// @Param	 		fts body request.
// @Accept 			json
// @Produce			json
// @Success 			200 {object}
func DeleteFeatureToggleHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusNoContent, gin.H{
		"msg": "DELETE",
	})

}
