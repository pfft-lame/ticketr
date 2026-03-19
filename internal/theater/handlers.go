package theater

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

func (h *handler) CreateTheater(c *echo.Context) error {
	var t createTheaterReq
	if err := c.Bind(&t); err != nil {
		return err
	}

	if err := c.Validate(t); err != nil {
		return err
	}

	res, err := h.s.CreateTheater(c.Request().Context(), t)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Theater created successfully!",
		Body:       res,
	})
}

func (h *handler) DeleteTheater(c *echo.Context) error {
	id := c.Param("id")

	err := h.s.DeleteTheaterById(c.Request().Context(), id)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Theater deleted successfully!",
	})
}

func (h *handler) UpdateTheater(c *echo.Context) error {
	id := c.Param("id")

	var t updateTheaterReq
	if err := c.Bind(&t); err != nil {
		return err
	}

	if err := c.Validate(t); err != nil {
		return err
	}

	res, err := h.s.UpdateTheaterById(c.Request().Context(), id, t)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Theater updated successfully!",
		Body:       res,
	})
}

func (h *handler) GetTheaterById(c *echo.Context) error {
	id := c.Param("id")

	res, err := h.s.GetTheaterById(c.Request().Context(), id)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Body:       res,
	})
}

func (h *handler) GetTheaters(c *echo.Context) error {
	cityId := c.QueryParam("city")
	if cityId == "" {
		res, err := h.s.GetAllTheaters(c.Request().Context())
		if err != nil {
			return apiresponse.DefaultServerError()
		}

		return c.JSON(http.StatusOK, apiresponse.ApiResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Body:       res,
		})
	}

	res, err := h.s.GetTheatersByCity(c.Request().Context(), cityId)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Body:       res,
	})
}
