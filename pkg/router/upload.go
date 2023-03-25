package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vesicash/upload-ms/pkg/controller/upload"
	"github.com/vesicash/upload-ms/utility"
)

func Upload(r *gin.Engine, ApiVersion string, validator *validator.Validate, logger *utility.Logger) *gin.Engine {
	upload := upload.Controller{Validator: validator, Logger: logger}

	authUrl := r.Group(fmt.Sprintf("%v", ApiVersion))
	{
		authUrl.POST("/files", upload.UploadFiles)
	}

	return r
}
