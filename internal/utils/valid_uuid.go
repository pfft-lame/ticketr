package utils

import (
	"net/http"

	apiresponse "ticketr/internal/api_response"

	"github.com/google/uuid"
)

func ValidUUID(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid UUID",
		}
	}

	return uid, nil
}
