package upload

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/vesicash/upload-ms/internal/config"
	"github.com/vesicash/upload-ms/internal/models"
	"github.com/vesicash/upload-ms/pkg/repository/storage/awss3"
	"github.com/vesicash/upload-ms/utility"
)

var (
	allowedMimeTypes = []string{"application/pdf",
		"image/png", "image/gif", "image/jpg", "text/plain", "image/jpeg",
		"video/mp4", "video/mpeg", "video/ogg", "video/quicktime",
		"application/msword", "application/doc", "application/docx", "application/pdf",
	}

	mg512 = 512
)

func UploadFilesService(c *gin.Context, logger *utility.Logger) ([]models.MultipleUploadResponse, int, error) {
	var (
		resp = []models.MultipleUploadResponse{}
	)

	form, err := c.MultipartForm()

	if err != nil {
		return resp, http.StatusBadRequest, fmt.Errorf("form data not provided")
	}

	files := form.File["files"]

	code, err := ValidateUploadRequest(c, logger)
	if err != nil {
		return resp, code, err
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return resp, http.StatusInternalServerError, err
		}

		defer file.Close()

		buff := make([]byte, mg512)

		_, err = file.Read(buff)
		if err != nil {
			return resp, http.StatusInternalServerError, err
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return resp, http.StatusInternalServerError, err
		}

		filename := fileHeader.Filename
		extension := filepath.Ext(fileHeader.Filename)
		fileMime := http.DetectContentType(buff)
		folder := ""
		id, _ := uuid.NewV4()
		shaFileName, err := utility.ShaHash(filename)
		if err != nil {
			return resp, http.StatusInternalServerError, err
		}

		fileMimeSlice := strings.Split(fileMime, "/")
		if fileMimeSlice[0] == "image" {
			folder = "images"
		} else {
			folder = "documents"
		}

		renamed := fmt.Sprintf("%v/%v/%v%v%v", config.GetConfig().App.Name, folder, shaFileName, id.String(), extension)
		url, err := awss3.UploadToS3(renamed, file)
		if err != nil {
			return resp, http.StatusInternalServerError, err
		}

		resp = append(resp, models.MultipleUploadResponse{
			OriginalName: filename,
			FileURL:      url,
		})
	}

	return resp, http.StatusOK, nil
}

func ValidateUploadRequest(c *gin.Context, logger *utility.Logger) (int, error) {
	var (
		form, _       = c.MultipartForm()
		files         = form.File["files"]
		maxSize       = 5048576
		maxSizeString = fmt.Sprintf("%vMB", int(math.Floor(float64(maxSize)/1000000)))
	)

	if len(files) < 1 {
		return http.StatusBadRequest, fmt.Errorf("At-least one file is required for upload.")
	}

	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			return http.StatusInternalServerError, err
		}

		defer file.Close()

		buff := make([]byte, mg512)

		_, err = file.Read(buff)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		filetype := http.DetectContentType(buff)
		if !contains(filetype, allowedMimeTypes) {
			return http.StatusBadRequest, fmt.Errorf("File type %v is not allowed", filetype)
		}

		_, err = file.Seek(0, io.SeekStart)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		fileSize := fileHeader.Size
		if fileSize > int64(maxSize) {
			return http.StatusBadRequest, fmt.Errorf("File size is greater than %v.", maxSizeString)
		}
	}

	return http.StatusOK, nil
}

func contains(v string, a []string) bool {
	for _, i := range a {
		if i == v {
			return true
		}
	}

	return false
}
