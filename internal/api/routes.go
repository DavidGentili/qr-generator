package api

import (
	"qr-generator/config"
	image_builder "qr-generator/internal/image"
	"qr-generator/internal/qr"

	"github.com/gin-gonic/gin"
)

func SetupRouter(buildInfo config.BuildInfo, qrBuilder *qr.QRBuilder, imageBuilder *image_builder.ImageBuilder) *gin.Engine {
	router := gin.New()
	group := router.Group("/api")
	handler := NewQrHandler(qrBuilder, imageBuilder)
	healthHandler := NewHealthHandler(buildInfo.Version, buildInfo.Commit, buildInfo.BuildDate)

	group.POST("/qr/generate", handler.GenerateQR)
	group.GET("/health", healthHandler.Health)

	return router
}
