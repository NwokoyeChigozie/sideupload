package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vesicash/upload-ms/pkg/controller/health"
	"github.com/vesicash/upload-ms/utility"
)

func Health(r *gin.Engine, ApiVersion string, validator *validator.Validate, logger *utility.Logger) *gin.Engine {
	health := health.Controller{Validator: validator, Logger: logger}

	healthUrl := r.Group(fmt.Sprintf("%v/upload", ApiVersion))
	{
		healthUrl.POST("/health", health.Post)
		healthUrl.GET("/health", health.Get)
	}
	return r
}
