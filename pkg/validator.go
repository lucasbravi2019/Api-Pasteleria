package pkg

import (
	"log"

	"github.com/go-playground/validator/v10"
)

func Validate(obj interface{}) error {
	validationErrors := validator.New().Struct(obj)
	if validationErrors != nil {
		log.Println("Error de validacion\n" + validationErrors.Error())
		return validationErrors
	}
	return nil
}
