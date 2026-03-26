package cities

import (
	"net/http"

	apiresponse "ticketr/internal/api_response"

	"github.com/labstack/echo/v5"
)

type handler struct {
	s Service
}

func NewHandler(s Service) *handler {
	return &handler{s}
}

func (h *handler) CreateCity(c *echo.Context) error {
	var city createCityReq
	if err := c.Bind(&city); err != nil {
		return err
	}

	if err := c.Validate(city); err != nil {
		return err
	}

	res, err := h.s.CreateCity(c.Request().Context(), city)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "City created!",
		Body:       res,
	})
}

func (h *handler) DeleteCity(c *echo.Context) error {
	id := c.Param("id")

	err := h.s.DeleteCity(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "City deleted Successfully!",
	})
}

func (h *handler) GetCity(c *echo.Context) error {
	id := c.Param("id")

	res, err := h.s.GetCityById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Body:       res,
	})
}

func (h *handler) GetAllCities(c *echo.Context) error {
	res, err := h.s.GetAllCities(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Body:       res,
	})
}
