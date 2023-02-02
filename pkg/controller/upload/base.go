package upload

import (
	"github.com/go-playground/validator/v10"
	"github.com/vesicash/upload-ms/utility"
)

type Controller struct {
	Validator *validator.Validate
	Logger    *utility.Logger
}
