package config

import (
	"os"
	image_builder "qr-generator/internal/image"
)

type Config struct {
	Port        string
	ImageParams image_builder.StyleImageParams
}

func LoadConfig() (*Config, error) {

	imageParams := image_builder.StyleImageParams{
		QRColor:         readEnv("QR_COLOR", "#000000"),
		ColorBackground: readEnv("COLOR_BACKGROUND", "#FFFFFF"),
		ColorBorder:     readEnv("COLOR_BORDER", "#000000"),
		ColorTitle:      readEnv("COLOR_TITLE", "#000000"),
		ColorSubtitle:   readEnv("COLOR_SUBTITLE", "#000000"),
		ColorMessage:    readEnv("COLOR_MESSAGE", "#000000"),
	}

	return &Config{
		Port:        readEnv("PORT", ":8080"),
		ImageParams: imageParams,
	}, nil
}

func readEnv(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
