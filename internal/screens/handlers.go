package screens

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

func (h *handler) CreateScreen(c *echo.Context) error {
	var screen createScreenReq
	if err := c.Bind(&screen); err != nil {
		return err
	}

	if err := c.Validate(screen); err != nil {
		return err
	}

	res, err := h.s.CreateScreen(c.Request().Context(), screen)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "Screen created successfully!",
		Body:       res,
	})
}

func (h *handler) UpdateScreen(c *echo.Context) error {
	id := c.Param("id")

	var screen updateScreenReq
	if err := c.Bind(&screen); err != nil {
		return err
	}

	if err := c.Validate(screen); err != nil {
		return err
	}

	res, err := h.s.UpdateScreenId(c.Request().Context(), id, screen)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "Screen updated successfully!",
		Body:       res,
	})
}

func (h *handler) GetScreenById(c *echo.Context) error {
	id := c.Param("id")

	res, err := h.s.GetScreenById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Body:       res,
	})
}

func (h *handler) DeleteScreenById(c *echo.Context) error {
	id := c.Param("id")

	err := h.s.DeleteScreenById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Screen deleted successfully!",
	})
}

func (h *handler) GetScreens(c *echo.Context) error {
	theaterId := c.QueryParam("theaterId")

	if theaterId == "" {
		res, err := h.s.GetAllScreens(c.Request().Context())
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, apiresponse.ApiResponse{
			StatusCode: http.StatusOK,
			Success:    true,
			Body:       res,
		})
	}

	res, err := h.s.GetScreenByTheaterId(c.Request().Context(), theaterId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Body:       res,
	})
}
