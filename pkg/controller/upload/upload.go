package upload

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vesicash/upload-ms/services/upload"
	"github.com/vesicash/upload-ms/utility"
)

func (base *Controller) UploadFiles(c *gin.Context) {

	urls, code, err := upload.UploadFilesService(c, base.Logger)
	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", err.Error(), err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "successful", urls)
	c.JSON(http.StatusCreated, rd)

}
