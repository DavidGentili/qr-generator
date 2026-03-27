package api

import (
	"log"
	"net/http"
	"time"

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

type QRHandler struct {
	qrBuilder    *qr.QRBuilder
	imageBuilder *image_builder.ImageBuilder
}

func NewQrHandler(qrBuilder *qr.QRBuilder, imageBuilder *image_builder.ImageBuilder) *QRHandler {
	return &QRHandler{
		qrBuilder:    qrBuilder,
		imageBuilder: imageBuilder,
	}
}

func (h *QRHandler) GenerateQR(c *gin.Context) {

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

type HealthHandler struct {
	version   string
	commit    string
	buildDate string
	startTime time.Time
}

func NewHealthHandler(version string, commit string, buildDate string) *HealthHandler {
	return &HealthHandler{
		version:   version,
		commit:    commit,
		buildDate: buildDate,
		startTime: time.Now(),
	}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "healthy",
		"service":   "qr-generator",
		"version":   h.version,
		"commit":    h.commit,
		"buildDate": h.buildDate,
		"uptime":    time.Since(h.startTime).String(),
	})
}
