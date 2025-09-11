package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(data any) []string {
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	var errors []string
	for _, e := range err.(validator.ValidationErrors) {
		fieldName := e.Field()
		if jsonTag := e.StructField(); jsonTag != "" {
			fieldName = strings.ToLower(e.Field())
		}

		switch e.Tag() {
		case "required":
			errors = append(errors, fmt.Sprintf("%s is required", fieldName))
		case "email":
			errors = append(errors, fmt.Sprintf("%s must be a valid email", fieldName))
		case "min":
			errors = append(errors, fmt.Sprintf("%s must be at least %s characters", fieldName, e.Param()))
		case "oneof":
			errors = append(errors, fmt.Sprintf("%s must be one of: %s", fieldName, e.Param()))
		default:
			errors = append(errors, fmt.Sprintf("%s is invalid", fieldName))
		}
	}
	return errors
}
