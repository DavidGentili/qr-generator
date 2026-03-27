package main

import (
	"log"
	"qr-generator/config"
	"qr-generator/internal/api"
	image_builder "qr-generator/internal/image"
	"qr-generator/internal/observability"
	"qr-generator/internal/qr"
)

var (
	Version   = "dev"
	Commit    = "none"
	BuildDate = "unknown"
)

func main() {
	configData, err := config.LoadConfig()
	buildInfo := config.BuildInfo{Version: Version, Commit: Commit, BuildDate: BuildDate}
	if err != nil {
		log.Fatalf("[Main] Error al cargar la configuración: %v", err)
	}

	qrBuilder := qr.NewQRBuilder(configData.ImageSizeParams.QrSize, configData.ImageParams.QRColor, configData.ImageParams.ColorBackground)
	imageBuilder, err := image_builder.NewImageBuilder(configData.ImageParams, configData.ImageSizeParams)
	if err != nil {
		log.Fatalf("[Main] Error al inicializar el image builder: %v", err)
	}

	observability.StartMetricsServer(configData.ObservabilityParams)

	router := api.SetupRouter(buildInfo, qrBuilder, imageBuilder)
	router.Run(configData.Port)
}
