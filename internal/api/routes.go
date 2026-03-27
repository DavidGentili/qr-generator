package api

import (
	image_builder "qr-generator/internal/image"
	"qr-generator/internal/qr"

	"github.com/gin-gonic/gin"
)

func SetupRouter(qrBuilder *qr.QRBuilder, imageBuilder *image_builder.ImageBuilder) *gin.Engine {
	router := gin.New()
	group := router.Group("/api")
	handler := NewHandler(qrBuilder, imageBuilder)

	group.POST("/qr/generate", handler.GenerateQR)

	return router
}
