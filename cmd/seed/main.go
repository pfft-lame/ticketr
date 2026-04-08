package main

import (
	"context"
	"log"
	"os"
	"time"

	"ticketr/internal/db"
	"ticketr/internal/db/queries"

	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	dsn := os.Getenv("GOOSE_DBSTRING")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	dbInstance, err := db.New(ctx, dsn)
	if err != nil {
		log.Fatalf("DB connection failed: %s", err)
	}

	q := queries.New(dbInstance)

	cityIds := AddCities(q)
	movies := AddMovies(q)
	theaterIds := AddTheaters(q, cityIds)
	screenIds := AddScreens(q, theaterIds)
	AddShows(q, screenIds, movies)
}

func AddCities(q queries.Querier) []uuid.UUID {
	cityId := make([]uuid.UUID, 0, 5)

	for _, s := range cities() {
		r, err := q.CreateCity(context.Background(), s)
		if err != nil {
			panic(err)
		}

		cityId = append(cityId, r.ID)
	}

	return cityId
}

func AddMovies(q queries.Querier) []queries.CreateMovieRow {
	moviesRows := make([]queries.CreateMovieRow, 0, 10)

	for _, m := range movies() {
		r, err := q.CreateMovie(context.Background(), m)
		if err != nil {
			panic(err)
		}
		moviesRows = append(moviesRows, r)
	}

	return moviesRows
}

func AddTheaters(q queries.Querier, cityIDs []uuid.UUID) []uuid.UUID {
	theaterIds := make([]uuid.UUID, 0, 30)

	for _, t := range theaters(cityIDs) {
		r, err := q.CreateTheater(context.Background(), t)
		if err != nil {
			panic(err)
		}

		theaterIds = append(theaterIds, r.ID)
	}

	return theaterIds
}

func AddScreens(q queries.Querier, theaterIDs []uuid.UUID) []uuid.UUID {
	screenIds := []uuid.UUID{}

	for _, s := range screens(theaterIDs) {
		r, err := q.CreateScreen(context.Background(), s)
		if err != nil {
			panic(err)
		}

		screenIds = append(screenIds, r.ID)
	}

	return screenIds
}

func AddShows(q queries.Querier, screenIDs []uuid.UUID, movies []queries.CreateMovieRow) []uuid.UUID {
	showIds := []uuid.UUID{}
	for _, s := range shows(movies, screenIDs) {
		r, err := q.CreateShow(context.Background(), s)
		if err != nil {
			panic(err)
		}

		showIds = append(showIds, r.ID)
	}

	return showIds
}
