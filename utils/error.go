package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func ValidationError(err error) map[string]string {

	errors := make(map[string]string)

	for _, err := range err.(validator.ValidationErrors) {

		field := strings.ToLower(err.Field())

		switch err.Tag() {

		case "required":
			errors[field] = field + " wajib diisi"

		case "email":
			errors[field] = "format email tidak valid"

		case "min":
			errors[field] = field + " minimal karakter tidak terpenuhi"

		}
	}

	return errors
}