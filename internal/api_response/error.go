package apiresponse

import (
	"fmt"
	"net/http"
	"reflect"

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

	if e, ok := err.(ApiError); ok && !isEmptyBody(e.Body) {
		c.JSON(e.StatusCode, ApiResponse{
			Success:    false,
			StatusCode: e.StatusCode,
			Errors:     e.Body,
		})
		return
	}

	if e, ok := err.(*echo.HTTPError); ok {
		c.JSON(e.Code, ApiResponse{
			Success:    false,
			StatusCode: e.Code,
			Errors:     e.Error(),
		})
		return
	}

	if _, ok := err.(*echo.BindingError); ok {
		c.JSON(http.StatusBadRequest, ApiResponse{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Errors:     "Invalid request payload",
		})
		return
	}

	fmt.Println(err.Error())

	c.JSON(http.StatusInternalServerError, ApiResponse{
		Success:    false,
		StatusCode: http.StatusInternalServerError,
		Errors:     "Something went wrong!",
	})
}

func isEmptyBody(body any) bool {
	if body == nil {
		return true
	}

	v := reflect.ValueOf(body)

	switch v.Kind() {
	case reflect.Map, reflect.Array, reflect.Slice:
		return v.Len() == 0

	case reflect.String:
		return v.Len() == 0

	case reflect.Pointer, reflect.Interface:
		return v.IsNil()

	default:
		return false
	}
}
