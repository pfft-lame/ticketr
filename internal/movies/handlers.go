package movies

import (
	"net/http"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/middlewares"

	"github.com/google/uuid"
	"github.com/labstack/echo/v5"
)

type handler struct {
	s Service
}

func NewHandler(s Service) *handler {
	return &handler{s}
}

func (h *handler) CreateMovie(c *echo.Context) error {
	var m createMovieReq

	if err := c.Bind(&m); err != nil {
		return err
	}

	if err := c.Validate(m); err != nil {
		return err
	}

	res, err := h.s.CreateMovie(c.Request().Context(), m)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "Movie created successfully",
		Body:       res,
	})
}

func (h *handler) GetMovieById(c *echo.Context) error {
	id := c.Param("id")

	res, err := h.s.GetMovieById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Body:       res,
	})
}

func (h *handler) DeleteMovieById(c *echo.Context) error {
	id := c.Param("id")

	err := h.s.DeleteMovieById(c.Request().Context(), id)
	if err != nil {
		return err
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

	updatedMovie, err := h.s.UpdateMovieById(c.Request().Context(), m, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Movie updated successfully!",
		Body:       updatedMovie,
	})
}

func (h *handler) GetMovies(c *echo.Context) error {
	cityId, ok := c.Get(middlewares.CITY_ID).(uuid.UUID)
	if !ok {
		return apiresponse.CityIdError()
	}

	query := c.QueryParam("query")
	if query != "" {
		res, err := h.s.GetMoviesByName(c.Request().Context(), query, cityId)
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, apiresponse.ApiResponse{
			StatusCode: http.StatusOK,
			Success:    true,
			Body:       res,
		})
	}

	/*
		// TODO: don't allow the 'user' role to exec this
	*/
	res, err := h.s.GetAllMovies(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		StatusCode: http.StatusOK,
		Success:    true,
		Body:       res,
	})
}

func (h *handler) GetUpcomingMovies(c *echo.Context) error {
	cityId, ok := c.Get(middlewares.CITY_ID).(uuid.UUID)
	if !ok {
		return apiresponse.CityIdError()
	}

	res, err := h.s.GetUpcomingMovies(c.Request().Context(), cityId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Body:       res,
	})
}
