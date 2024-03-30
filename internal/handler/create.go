package toggle

import (
	"github.com/Ajaax2/feature-toggle-api/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createToggleResponse struct {
	ID string `json:"id"`
}

// CreateFeatureToggleHandler
// @Summary Create a new feature toggle.
// @Description This endpoint is responsible for creating a new feature toggle, saving on MongoDB.
// @Tags CREATE
// @Accept  json
// @Produce  json
// @Success 200 {object}  createToggleResponse
// @Router /api/v1/toggle [post]
func CreateFeatureToggleHandler(ctx *gin.Context) {

	v, err := model.NewVersion("123")
	toggle, err := model.NewFeatureToggle(v, "toggle", true)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createToggleResponse{ID: toggle.ID})
}
