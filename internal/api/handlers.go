package api

import (
	"net/http"

	image_builder "qr-generator/internal/image"
	"qr-generator/internal/qr"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	qrBuilder    *qr.QRBuilder
	imageBuilder *image_builder.ImageBuilder
}

func NewHandler(qrBuilder *qr.QRBuilder, imageBuilder *image_builder.ImageBuilder) *Handler {
	return &Handler{
		qrBuilder:    qrBuilder,
		imageBuilder: imageBuilder,
	}
}

func (h *Handler) GenerateQR(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "QR generated successfully",
	})
}
