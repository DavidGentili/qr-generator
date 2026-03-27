package api

import (
	"log"
	"net/http"

	image_builder "qr-generator/internal/image"
	"qr-generator/internal/qr"

	"github.com/gin-gonic/gin"
)

type GenerateQrBody struct {
	Link     string `json:"link" binding:"required"`
	Title    string `json:"title" binding:"required"`
	Subtitle string `json:"subtitle"`
	Message  string `json:"message"`
}

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

	var body GenerateQrBody
	if bodyErr := c.ShouldBindJSON(&body); bodyErr != nil {
		log.Printf("[Handler] Error binding body: %v", bodyErr)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Error binding body",
		})
		return
	}

	qrImage, qrError := h.qrBuilder.GenerateQR(body.Link)
	if qrError != nil {
		log.Printf("[Handler] Error generating QR: %v", qrError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating QR",
		})
		return
	}

	imageBuffer, imageError := h.imageBuilder.BuildImage(qrImage, image_builder.ImageData{
		Title:    body.Title,
		Subtitle: body.Subtitle,
		Message:  body.Message,
	})
	if imageError != nil {
		log.Printf("[Handler] Error building image: %v", imageError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error building image",
		})
		return
	}

	c.Data(http.StatusOK, "image/png", imageBuffer)
}
