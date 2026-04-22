package apiresponse

import "net/http"

func InvalidUUID() error {
	return ApiError{
		StatusCode: http.StatusBadRequest,
		Body:       "Invalid UUID",
	}
}

func CityIdError() error {
	return ApiError{
		StatusCode: http.StatusBadRequest,
		Body:       "X-City-Id missing or invalid.",
	}
}
