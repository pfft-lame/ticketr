package apiresponse

import "net/http"

func DefaultServerError() error {
	return ApiError{
		StatusCode: http.StatusInternalServerError,
		Body:       "Something went wrong!",
	}
}

func InvalidRequestError() error {
	return ApiError{
		StatusCode: http.StatusBadRequest,
		Body:       "Invalid request payload",
	}
}

func InvalidUUID() error {
	return ApiError{
		StatusCode: http.StatusBadRequest,
		Body:       "Invalid UUID",
	}
}
