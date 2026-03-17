package main

import (
	"net/http"
	"time"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/movie"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func (app *application) mount() http.Handler {
	e := echo.New()
	e.Validator = &responseValidator{app.validate}
	e.HTTPErrorHandler = apiresponse.GlobalErrorResponse

	// e.Use(middleware.RequestLogger())
	e.Use(middleware.RemoveTrailingSlash())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hi there")
	})

	api := e.Group("/api/v1")

	movieService := movie.NewService(app.queries)
	movieHandler := movie.NewHandler(movieService)
	movies := api.Group("/movies")
	movies.POST("", movieHandler.CreateMovie)
	movies.GET("/:id", movieHandler.GetMovieById)
	movies.DELETE("/:id", movieHandler.DeleteMovieById)
	movies.PATCH("/:id", movieHandler.UpdateMovieById)

	return e
}

func (app *application) server(h http.Handler) *http.Server {
	svr := &http.Server{
		Addr:         ":" + app.cfg.port,
		Handler:      h,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
	}

	return svr
}
