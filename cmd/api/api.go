package main

import (
	"net/http"
	"time"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/cities"
	"ticketr/internal/middlewares"
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
	public := api.Group("")
	cityPublic := api.Group("", middlewares.CityContextMiddleware)

	// Movies
	movieService := movies.NewService(app.queries)
	movieHandler := movies.NewHandler(movieService)

	public.GET("/movies/:id", movieHandler.GetMovieById)
	public.POST("/movies", movieHandler.CreateMovie)
	public.DELETE("/movies/:id", movieHandler.DeleteMovieById)
	public.PATCH("/movies/:id", movieHandler.UpdateMovieById)
	cityPublic.GET("/movies/upcoming", movieHandler.GetUpcomingMovies)
	cityPublic.GET("/movies", movieHandler.GetMovies)

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

	public.POST("/theaters", theaterHandler.CreateTheater)
	cityPublic.GET("/theaters", theaterHandler.GetTheaters)
	public.GET("/theaters/:id", theaterHandler.GetTheaterById)
	public.PATCH("/theaters/:id", theaterHandler.UpdateTheater)
	public.DELETE("/theaters/:id", theaterHandler.DeleteTheater)
	cityPublic.GET("/theaters/:id/upcoming", theaterHandler.GetUpcomingMoviesInTheater)

	// screens
	screensService := screens.NewService(app.queries)
	screensHandler := screens.NewHandler(screensService)

	public.POST("/screens", screensHandler.CreateScreen)
	public.GET("/screens/:id", screensHandler.GetScreenById)
	public.PATCH("/screens/:id", screensHandler.UpdateScreen)
	public.DELETE("/screens/:id", screensHandler.DeleteScreenById)
	public.GET("/screens", screensHandler.GetScreens)

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
