package handler

import config "github.com/Ajaax2/feature-toggle-api/configs"

var (
	logger *config.Logger
)

func Init() {
	logger = config.GetLogger("handler")
}
