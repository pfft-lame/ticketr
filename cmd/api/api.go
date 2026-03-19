package main

import (
	"net/http"
	"time"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/city"
	"ticketr/internal/movie"
	"ticketr/internal/theater"

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
	movieService := movie.NewService(app.queries)
	movieHandler := movie.NewHandler(movieService)

	movies := api.Group("/movies")
	movies.POST("", movieHandler.CreateMovie)
	movies.GET("/:id", movieHandler.GetMovieById)
	movies.DELETE("/:id", movieHandler.DeleteMovieById)
	movies.PATCH("/:id", movieHandler.UpdateMovieById)

	// cities
	cityService := city.NewService(app.queries)
	cityHandler := city.NewHandler(cityService)

	city := api.Group("/cities")
	city.POST("", cityHandler.CreateCity)
	city.GET("", cityHandler.GetAllCities)
	city.GET("/:id", cityHandler.GetCity)
	city.DELETE("/:id", cityHandler.DeleteCity)

	// Theaters
	theaterService := theater.NewService(app.queries)
	theaterHandler := theater.NewHandler(theaterService)

	theater := api.Group("/theaters")
	theater.POST("", theaterHandler.CreateTheater)
	theater.GET("", theaterHandler.GetTheaters)
	theater.GET("/:id", theaterHandler.GetTheaterById)
	theater.PATCH("/:id", theaterHandler.UpdateTheater)
	theater.DELETE("/:id", theaterHandler.DeleteTheater)

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
