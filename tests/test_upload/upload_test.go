package test_upload

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vesicash/upload-ms/pkg/controller/upload"
	"github.com/vesicash/upload-ms/pkg/repository/storage/awss3"
	tst "github.com/vesicash/upload-ms/tests"
)

func TestUploadFiles(t *testing.T) {
	logger := tst.Setup()
	gin.SetMode(gin.TestMode)
	validatorRef := validator.New()

	upload := upload.Controller{Validator: validatorRef, Logger: logger}
	r := gin.Default()

	tests := []struct {
		Name         string
		RequestBody  bool
		ExpectedCode int
		Headers      map[string]string
		Message      string
	}{
		{
			Name:         "ok upload",
			RequestBody:  true,
			ExpectedCode: http.StatusCreated,
			Message:      "successful",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}, {
			Name:         "no file",
			RequestBody:  false,
			ExpectedCode: http.StatusBadRequest,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		},
	}

	uploadUrl := r.Group(fmt.Sprintf("%v", "v2"))
	{
		uploadUrl.POST("/files", upload.UploadFiles)

	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var payload bytes.Buffer
			writer := multipart.NewWriter(&payload)
			// json.NewEncoder(&b).Encode(test.RequestBody)
			URI := url.URL{Path: "/v2/files"}
			if test.RequestBody {
				file, errFile1 := os.Open("./test_image.png")
				if errFile1 != nil {
					panic("errFile1" + errFile1.Error())
				}

				defer file.Close()
				part1,
					errFile1 := writer.CreateFormFile("files", filepath.Base("./test_image.png"))
				if errFile1 != nil {
					panic("errFile1" + errFile1.Error())
				}

				_, errFile1 = io.Copy(part1, file)
				if errFile1 != nil {
					panic("errFile1" + errFile1.Error())
				}

			}

			err := writer.Close()
			if err != nil {
				panic("writer err: " + err.Error())
			}

			req, err := http.NewRequest(http.MethodPost, URI.String(), &payload)
			if err != nil {
				t.Fatal(err)
			}

			for i, v := range test.Headers {
				req.Header.Set(i, v)
			}

			req.Header.Set("Content-Type", writer.FormDataContentType())

			rr := httptest.NewRecorder()
			r.ServeHTTP(rr, req)

			objectName := awss3.DeleteObjectInTestMode
			if objectName != "" {
				awss3.DeleteUpload(objectName)
			}

			tst.AssertStatusCode(t, rr.Code, test.ExpectedCode)

			data := tst.ParseResponse(rr)

			code := int(data["code"].(float64))
			tst.AssertStatusCode(t, code, test.ExpectedCode)

			if test.Message != "" {
				message := data["message"]
				if message != nil {
					tst.AssertResponseMessage(t, message.(string), test.Message)
				} else {
					tst.AssertResponseMessage(t, "", test.Message)
				}

			}

		})

	}

}
