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
		QRColor:          readEnv("QR_COLOR", "#000000"),
		ColorBackground:  readEnv("COLOR_BACKGROUND", "#FFFFFF"),
		ColorBorder:      readEnv("COLOR_BORDER", "#000000"),
		ColorTitle:       readEnv("COLOR_TITLE", "#000000"),
		ColorSubtitle:    readEnv("COLOR_SUBTITLE", "#000000"),
		ColorMessage:     readEnv("COLOR_MESSAGE", "#000000"),
		FontTitleSize:    readEnvFloat("FONT_TITLE_SIZE", 28),
		FontSubtitleSize: readEnvFloat("FONT_SUBTITLE_SIZE", 24),
		FontMessageSize:  readEnvFloat("FONT_MESSAGE_SIZE", 24),
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

func readEnvFloat(key string, defaultValue float64) float64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}
	return floatValue
}
