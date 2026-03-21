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

func GlobalErrorResponse(c *echo.Context, err error) {
	if resp, uErr := echo.UnwrapResponse(c.Response()); uErr == nil {
		if resp.Committed {
			return
		}
	}

	apiErr, ok := err.(ApiError)
	if ok {
		code := apiErr.StatusCode
		var body any
		if isEmptyBody(apiErr.Body) {
			body = "Something went wrong!"
			code = http.StatusInternalServerError
		} else {
			body = apiErr.Body
		}

		c.JSON(apiErr.StatusCode, ApiResponse{
			Success:    false,
			StatusCode: code,
			Errors:     body,
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

func isEmptyBody(body any) bool {
	if body == nil {
		return true
	}

	switch v := body.(type) {
	case string:
		return len(v) == 0
	case map[string]string:
		return len(v) == 0
	case []string:
		return len(v) == 0
	default:
		return false
	}
}
