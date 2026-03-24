package image_builder

import "image"

type StyleImageParams struct {
	QRColor         string
	ColorBackground string
	ColorBorder     string
	ColorTitle      string
	ColorSubtitle   string
	ColorMessage    string
}

type ImageData struct {
	Title    string
	Subtitle string
	Message  string
}

type ImageBuilder struct {
	Params StyleImageParams
}

func NewImageBuilder(params StyleImageParams) *ImageBuilder {
	return &ImageBuilder{Params: params}
}

func (b *ImageBuilder) BuildImage(qr *image.Image, data ImageData) (*image.Image, error) {
	return nil, nil
}
