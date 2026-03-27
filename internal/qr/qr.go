package qr

import (
	"image"
	"image/color"
	"log"

	colorUtil "qr-generator/internal/colorUtil"

	qrcode "github.com/skip2/go-qrcode"
)

type QRBuilder struct {
	QrSize          int
	QRColor         string
	BackgroundColor string
}

func NewQRBuilder(qrSize int, qrColor string, backgroundColor string) *QRBuilder {
	return &QRBuilder{QrSize: qrSize, QRColor: qrColor, BackgroundColor: backgroundColor}
}

func (b *QRBuilder) GenerateQR(link string) (*image.Image, error) {

	qr, constructorError := qrcode.New(link, qrcode.Medium)

	if constructorError != nil {
		return nil, constructorError
	}

	qr.BackgroundColor = b.parseHexColor(b.BackgroundColor)
	qr.ForegroundColor = b.parseHexColor(b.QRColor)
	qr.VersionNumber = 10
	qr.DisableBorder = true
	writeError := qr.WriteFile(b.QrSize, "./tmp/qr/qr.png") // Se escribe en fs para probar en local
	if writeError != nil {
		return nil, writeError
	}
	qrImage := qr.Image(b.QrSize)

	return &qrImage, nil
}

func (b *QRBuilder) parseHexColor(hexColor string) color.Color {
	parsed, err := colorUtil.ParseHexColor(hexColor)
	if err != nil {
		log.Printf("[QRBuilder] Error parsing hex color: %v", err)
		return color.Black
	}
	return parsed
}
