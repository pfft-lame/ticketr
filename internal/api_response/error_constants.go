package apiresponse

import "net/http"

func InvalidUUID() error {
	return ApiError{
		StatusCode: http.StatusBadRequest,
		Body:       "Invalid UUID",
	}
}
