package main

import (
	"net/http"
	"time"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/cities"
	"ticketr/internal/movies"
	"ticketr/internal/screens"
	"ticketr/internal/theaters"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func (app *application) mount() http.Handler {
	e := echo.New()
	e.Validator = &responseValidator{app.validate}
	e.HTTPErrorHandler = apiresponse.GlobalErrorResponse

	e.Pre(middleware.RemoveTrailingSlash())

	// e.Use(middleware.RequestLogger())

	e.GET("/", func(c *echo.Context) error {
		return c.String(http.StatusOK, "Hi there")
	})

	api := e.Group("/api/v1")

	// Movies
	movieService := movies.NewService(app.queries)
	movieHandler := movies.NewHandler(movieService)

	movies := api.Group("/movies")
	movies.POST("", movieHandler.CreateMovie)
	movies.GET("/:id", movieHandler.GetMovieById)
	movies.DELETE("/:id", movieHandler.DeleteMovieById)
	movies.PATCH("/:id", movieHandler.UpdateMovieById)
	movies.GET("", movieHandler.GetMovieByName)

	// cities
	cityService := cities.NewService(app.queries)
	cityHandler := cities.NewHandler(cityService)

	cities := api.Group("/cities")
	cities.POST("", cityHandler.CreateCity)
	cities.GET("", cityHandler.GetAllCities)
	cities.GET("/:id", cityHandler.GetCity)
	cities.DELETE("/:id", cityHandler.DeleteCity)

	// Theaters
	theaterService := theaters.NewService(app.queries)
	theaterHandler := theaters.NewHandler(theaterService)

	theater := api.Group("/theaters")
	theater.POST("", theaterHandler.CreateTheater)
	theater.GET("", theaterHandler.GetTheaters)
	theater.GET("/:id", theaterHandler.GetTheaterById)
	theater.PATCH("/:id", theaterHandler.UpdateTheater)
	theater.DELETE("/:id", theaterHandler.DeleteTheater)

	// screens
	screensService := screens.NewService(app.queries)
	screensHandler := screens.NewHandler(screensService)

	screens := api.Group("/screens")
	screens.POST("", screensHandler.CreateScreen)
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
