package handler

import (
	"net/http"

	"github.com/Ajaax2/feature-toggle-api/internal/model"
	"github.com/gin-gonic/gin"
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
	version, err := model.NewVersion("123")
	ft, err := model.NewFeatureToggle(version, "toggle", true)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, createToggleResponse{ID: ft.ID})
}
