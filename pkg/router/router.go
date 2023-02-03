package router

import (
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vesicash/upload-ms/pkg/middleware"
	"github.com/vesicash/upload-ms/utility"
)

func Setup(logger *utility.Logger, validator *validator.Validate) *gin.Engine {
	r := gin.New()

	// Middlewares
	// r.Use(gin.Logger())
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.CORS())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.MaxMultipartMemory = 1 << 20 // 1MB

	ApiVersion := "v2"
	Health(r, ApiVersion, validator, logger)
	Upload(r, ApiVersion, validator, logger)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    404,
			"status":  http.StatusNotFound,
		})
	})

	return r
}