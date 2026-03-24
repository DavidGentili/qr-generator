package api

import (
	image_builder "qr-generator/internal/image"
	"qr-generator/internal/qr"

	"github.com/gin-gonic/gin"
)

func SetupRouter(qrBuilder *qr.QRBuilder, imageBuilder *image_builder.ImageBuilder) *gin.Engine {
	router := gin.Default()
	handler := NewHandler(qrBuilder, imageBuilder)

	router.POST("/qr/generate", handler.GenerateQR)

	return router
}
