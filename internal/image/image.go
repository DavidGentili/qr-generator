package image_builder

import (
	"bytes"
	"image"
	"image/png"

	"gopkg.in/fogleman/gg.v1"
)

type StyleImageParams struct {
	QRColor         string
	ColorBackground string
	ColorBorder     string
	ColorTitle      string
	ColorSubtitle   string
	ColorMessage    string
}

type ImageSizeParams struct {
	QrSize    int
	QrPadding int
}

type ImageData struct {
	Title    string
	Subtitle string
	Message  string
}

type ImageBuilder struct {
	StyleParams StyleImageParams
	SizeParams  ImageSizeParams
}

func NewImageBuilder(params StyleImageParams, sizeParams ImageSizeParams) *ImageBuilder {
	return &ImageBuilder{StyleParams: params, SizeParams: sizeParams}
}

func (b *ImageBuilder) BuildImage(qr *image.Image, data ImageData) ([]byte, error) {
	body := b.generateBody(*qr)
	header := b.generateHeader(data.Title, data.Subtitle, body.Bounds().Dx())
	footer := b.generateFooter(data.Message, body.Bounds().Dx())

	compositeImage := b.compositeImage([]image.Image{*header, body, *footer})

	bgBuffer, bgBufferError := b.encodeImage(compositeImage)
	if bgBufferError != nil {
		return nil, bgBufferError
	}
	return bgBuffer, nil
}

func (b *ImageBuilder) encodeImage(image *image.Image) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	if err := png.Encode(buffer, *image); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (b *ImageBuilder) generateBody(image image.Image) image.Image {
	width := (image).Bounds().Dx()
	height := (image).Bounds().Dy()
	paddingComponent := b.generateBackground(width+b.SizeParams.QrPadding*2, height+b.SizeParams.QrPadding*2)

	paddingComponent.DrawImage(image, b.SizeParams.QrPadding, b.SizeParams.QrPadding)
	return paddingComponent.Image()
}

func (b *ImageBuilder) generateHeader(title string, subtitle string, width int) *image.Image {

	headerComponent := b.generateBackground(width, 64)

	headerComponent.SetHexColor(b.StyleParams.ColorTitle)
	b.writeText(headerComponent, title, width, 20)

	if subtitle != "" {
		headerComponent.SetHexColor(b.StyleParams.ColorSubtitle)
		b.writeText(headerComponent, subtitle, width, 40)
	}

	headerImage := headerComponent.Image()

	return &headerImage
}

func (b *ImageBuilder) generateFooter(message string, width int) *image.Image {
	footerComponent := b.generateBackground(width, 64)
	footerComponent.SetHexColor(b.StyleParams.ColorMessage)
	b.writeText(footerComponent, message, width, 20)
	footerImage := footerComponent.Image()
	return &footerImage
}

func (b *ImageBuilder) compositeImage(partialImages []image.Image) *image.Image {
	width := partialImages[0].Bounds().Dx()
	height := b.getHeight(partialImages)
	compositeComponent := gg.NewContext(width, height)
	y := 0
	for _, img := range partialImages {
		compositeComponent.DrawImage(img, 0, y)
		y += img.Bounds().Dy()
	}
	compositeImage := compositeComponent.Image()
	return &compositeImage

}

func (b *ImageBuilder) writeText(ctx *gg.Context, text string, width int, offsetY float64) {
	ctx.DrawStringWrapped(text, 0, offsetY, 0, 0, float64(width), 1.5, gg.AlignCenter)
}

func (b *ImageBuilder) generateBackground(width int, height int) *gg.Context {
	bgComponent := gg.NewContext(width, height)
	bgComponent.SetHexColor(b.StyleParams.ColorBackground)
	bgComponent.Clear()
	return bgComponent
}

func (b *ImageBuilder) getHeight(image []image.Image) int {
	height := 0
	for _, img := range image {
		height += img.Bounds().Dy()
	}
	return height
}
