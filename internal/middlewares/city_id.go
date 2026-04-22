package middlewares

import (
	"net/http"

	apiresponse "ticketr/internal/api_response"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

func CityContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		cityIdStr := c.Request().Header.Get("X-City-Id")

		if cityIdStr == "" {
			return apiresponse.ApiError{
				StatusCode: http.StatusBadRequest,
				Body:       "Missing X-City-Id header",
			}
		}

		cityId, err := uuid.Parse(cityIdStr)
		if err != nil {
			return apiresponse.ApiError{
				StatusCode: http.StatusBadRequest,
				Body:       "Invalid X-City-Id header.",
			}
		}

		c.Set(CITY_ID, cityId)

		return next(c)
	}
}
