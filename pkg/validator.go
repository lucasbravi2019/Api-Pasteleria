package pkg

import (
	"github.com/go-playground/validator/v10"
)

func Validate(obj interface{}) error {
	validationErrors := validator.New().Struct(obj)
	if validationErrors != nil {
		return validationErrors
	}
	return nil
}
