package config

import (
	"os"
	image_builder "qr-generator/internal/image"
	"strconv"
)

type Config struct {
	Port            string
	ImageParams     image_builder.StyleImageParams
	ImageSizeParams image_builder.ImageSizeParams
}

func LoadConfig() (*Config, error) {

	imageParams := image_builder.StyleImageParams{
		QRColor:         readEnv("QR_COLOR", "#000000"),
		ColorBackground: readEnv("COLOR_BACKGROUND", "#FCAFFD"),
		ColorBorder:     readEnv("COLOR_BORDER", "#000000"),
		ColorTitle:      readEnv("COLOR_TITLE", "#000000"),
		ColorSubtitle:   readEnv("COLOR_SUBTITLE", "#000000"),
		ColorMessage:    readEnv("COLOR_MESSAGE", "#000000"),
	}

	imageSizeParams := image_builder.ImageSizeParams{
		QrSize:    readEnvInt("QR_SIZE", 256),
		QrPadding: readEnvInt("QR_PADDING", 32),
	}

	return &Config{
		Port:            readEnv("PORT", ":8080"),
		ImageParams:     imageParams,
		ImageSizeParams: imageSizeParams,
	}, nil
}

func readEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func readEnvInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}
	return intValue
}
