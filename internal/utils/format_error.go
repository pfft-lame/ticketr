package utils

import (
	"strings"

	"github.com/go-playground/validator/v10"
)

func FormatValidationErrs(e error) map[string]string {
	errs := make(map[string]string)

	err, ok := e.(validator.ValidationErrors)
	if !ok {
		return nil
	}

	for _, e := range err {
		f := e.Field()

		switch e.Tag() {
		case "required":
			errs[f] = f + " is required"
		case "email":
			errs[f] = "Invalid email address"
		case "dive":
			errs[f] = f + " must be a list"
		case "min":
			errs[f] = f + " must be at least the size of " + e.Param()
		case "max":
			errs[f] = f + " should not exceed the size " + e.Param()
		case "oneof":
			vals := strings.Split(e.Param(), " ")
			valStr := strings.Join(vals, ", ")
			errs[f] = f + " can have following values: " + valStr
		default:
			errs[f] = "Invalid value"
		}
	}

	return errs
}
