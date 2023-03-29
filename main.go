package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator/v10"
	"github.com/vesicash/upload-ms/internal/config"

	"github.com/vesicash/upload-ms/utility"

	"github.com/vesicash/upload-ms/pkg/repository/storage/awss3"
	"github.com/vesicash/upload-ms/pkg/router"
)

func main() {
	logger := utility.NewLogger() //Warning !!!!! Do not recreate this action anywhere on the app

	configuration := config.Setup(logger, "./app")
	awss3.ConnectAws(logger)

	validatorRef := validator.New()
	r := router.Setup(logger, validatorRef, &configuration.App)
	rM := router.SetupMetrics(&configuration.App)

	go func(logger *utility.Logger, metricsPort string) {
		utility.LogAndPrint(logger, fmt.Sprintf("Metric Server is starting at 127.0.0.1:%s", metricsPort))
		err := rM.Run(":" + metricsPort)
		if err != nil {
			log.Fatal(err)
		}
	}(logger, configuration.Server.MetricsPort)

	utility.LogAndPrint(logger, fmt.Sprintf("Server is starting at 127.0.0.1:%s", configuration.Server.Port))
	log.Fatal(r.Run(":" + configuration.Server.Port))
}
