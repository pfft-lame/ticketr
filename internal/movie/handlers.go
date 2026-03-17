package movie

import (
	"fmt"
	"net/http"

	apiresponse "ticketr/internal/api_response"

	"github.com/labstack/echo/v5"
)

type handler struct {
	service Service
}

func NewHandler(s Service) *handler {
	return &handler{s}
}

func (h *handler) CreateMovie(c *echo.Context) error {
	var m createMovieReq

	if err := c.Bind(&m); err != nil {
		fmt.Println("create movie req err")
		return err
	}

	if err := c.Validate(m); err != nil {
		return err
	}

	id, err := h.service.CreateMovie(c.Request().Context(), m)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "Movie created successfully",
		Body: map[string]string{
			"movie_id": id.String(),
		},
	})
}

func (h *handler) GetMovieById(c *echo.Context) error {
	id := c.Param("id")

	res, err := h.service.GetMovieById(c.Request().Context(), id)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Body:       res,
	})
}

func (h *handler) DeleteMovieById(c *echo.Context) error {
	id := c.Param("id")

	err := h.service.DeleteMovieById(c.Request().Context(), id)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Body:       "Movie deleted successfully!",
	})
}

func (h *handler) UpdateMovieById(c *echo.Context) error {
	id := c.Param("id")

	var m updateMovieReq
	if err := c.Bind(&m); err != nil {
		return err
	}

	if err := c.Validate(m); err != nil {
		return err
	}

	updatedMovie, err := h.service.UpdateMovieById(c.Request().Context(), m, id)
	if err != nil {
		if _, ok := err.(apiresponse.ApiError); ok {
			return err
		}

		return apiresponse.DefaultServerError()
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Movie updated successfully!",
		Body:       updatedMovie,
	})
}
