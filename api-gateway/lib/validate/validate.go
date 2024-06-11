package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ErrorResponse struct {
	Status  string   `json:"status"`
	Message []string `json:"message"`
}

func FormatValidationError(err error) ErrorResponse {
	var errors []string
	if validationErr, ok := err.(validator.ValidationErrors); ok {
		for _, v := range validationErr {
			switch v.Tag() {
			case "required":
				errors = append(errors, fmt.Sprintf("%s is required", v.Field()))
			case "email":
				errors = append(errors, fmt.Sprintf("%s is not valid email", v.Field()))
			case "gte":
				errors = append(errors, fmt.Sprintf("%s value must be greater than %s", v.Field(), v.Param()))
			case "lte":
				errors = append(errors, fmt.Sprintf("%s value must be lower than %s", v.Field(), v.Param()))
			case "min":
				errors = append(errors, fmt.Sprintf("%s character must be min %s", v.Field(), v.Param()))
			case "max":
				errors = append(errors, fmt.Sprintf("%s character must be max %s", v.Field(), v.Param()))
			default:
				errors = append(errors, v.Error())
			}
		}
	} else {
		errors = append(errors, err.Error())
	}

	response := ErrorResponse{
		Status:  "error",
		Message: errors,
	}

	return response
}
