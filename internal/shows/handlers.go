package shows

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

func (h *handler) CreateShow(c *echo.Context) error {
	var show createShowReq
	if err := c.Bind(&show); err != nil {
		return err
	}

	if err := c.Validate(show); err != nil {
		return err
	}

	res, err := h.s.CreateShow(c.Request().Context(), show)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "Show created successfully!",
		Body:       res,
	})
}

func (h *handler) UpdateShow(c *echo.Context) error {
	var show updateShowReq
	if err := c.Bind(&show); err != nil {
		return err
	}

	id := c.Param("id")

	res, err := h.s.UpdateShowById(c.Request().Context(), show, id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Show updated successfully!",
		Body:       res,
	})
}

func (h *handler) DeleteShow(c *echo.Context) error {
	id := c.Param("id")

	err := h.s.DeleteShowById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "Show deleted successfully!",
	})
}

func (h *handler) GetShowsById(c *echo.Context) error {
	id := c.Param("id")

	res, err := h.s.GetShowInfoById(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Body:       res,
	})
}

func (h *handler) GetShows(c *echo.Context) error {
	cityId, ok := c.Get(middlewares.CITY_ID).(uuid.UUID)
	if !ok {
		return apiresponse.CityIdError()
	}

	movieId := c.QueryParam("movie_id")
	if movieId != "" {
		res, err := h.s.GetShowsByMovieId(c.Request().Context(), movieId, cityId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, apiresponse.ApiResponse{
			StatusCode: http.StatusOK,
			Success:    true,
			Body:       res,
		})
	}

	theaterId := c.QueryParam("theater_id")
	if theaterId != "" {
		res, err := h.s.GetShowsByTheaterId(c.Request().Context(), movieId)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, apiresponse.ApiResponse{
			StatusCode: http.StatusOK,
			Success:    true,
			Body:       res,
		})

	}

	res, err := h.s.GetShowsByCityId(c.Request().Context(), cityId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, apiresponse.ApiResponse{
		Success:    true,
		StatusCode: http.StatusOK,
		Body:       res,
	})
}
