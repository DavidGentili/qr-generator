package qr

import "image"

type QRBuilder struct {
}

func NewQRBuilder() *QRBuilder {
	return &QRBuilder{}
}

func (b *QRBuilder) GenerateQR(link string) (*image.Image, error) {
	return nil, nil
}
