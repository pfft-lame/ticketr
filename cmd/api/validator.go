package main

import (
	"net/http"
	"reflect"
	"strings"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/utils"

	"github.com/go-playground/validator/v10"
)

func newValidator() *validator.Validate {
	v := validator.New()

	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	return v
}

type responseValidator struct {
	v *validator.Validate
}

func (r *responseValidator) Validate(s any) error {
	err := r.v.Struct(s)
	if err != nil {
		return apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       utils.FormatValidationErrs(err),
		}
	}

	return nil
}
