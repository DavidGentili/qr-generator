package main

import (
	"log"
	"qr-generator/config"
	"qr-generator/internal/api"
	image_builder "qr-generator/internal/image"
	"qr-generator/internal/qr"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[Main] Error al cargar la configuración: %v", err)
	}

	qrBuilder := qr.NewQRBuilder()
	imageBuilder := image_builder.NewImageBuilder(config.ImageParams)

	router := api.SetupRouter(qrBuilder, imageBuilder)
	router.Run(config.Port)
}
