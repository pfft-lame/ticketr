package apiresponse

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

type ApiError struct {
	StatusCode int
	Body       any
}

func (e ApiError) Error() string {
	return fmt.Sprint(e.Body)
}

func DefaultServerError() error {
	return ApiError{
		StatusCode: http.StatusInternalServerError,
		Body:       "Something went wrong!",
	}
}

func GlobalErrorResponse(c *echo.Context, err error) {
	if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
		if resp.Committed {
			return
		}
	}

	apiErr, ok := err.(ApiError)
	if ok {
		c.JSON(apiErr.StatusCode, ApiResponse{
			Success:    false,
			StatusCode: apiErr.StatusCode,
			Errors:     apiErr.Body,
		})
		return
	}

	if echoErr, ok := err.(*echo.HTTPError); ok {
		c.JSON(echoErr.Code, ApiResponse{
			Success:    false,
			StatusCode: echoErr.Code,
			Errors:     echoErr.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, ApiResponse{
		Success:    false,
		StatusCode: http.StatusInternalServerError,
		Errors:     err.Error(),
	})
}
