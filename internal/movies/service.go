package movies

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/db"
	repo "ticketr/internal/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	CreateMovie(ctx context.Context, movie createMovieReq) (repo.CreateMovieRow, error)
	GetMovieById(ctx context.Context, id string) (repo.GetMovieByIdRow, error)
	DeleteMovieById(ctx context.Context, id string) error
	UpdateMovieById(ctx context.Context, movie updateMovieReq, id string) (repo.UpdateMovieByIdRow, error)
	GetMoviesByName(ctx context.Context, query string, cityId uuid.UUID) ([]repo.GetMoviesByNameRow, error)
	GetUpcomingMovies(ctx context.Context, cityId uuid.UUID) ([]repo.GetUpcomingMoviesRow, error)
	GetAllMovies(ctx context.Context) ([]repo.GetAllMoviesRow, error)
}

type svc struct {
	repo repo.Querier
}

func NewService(q repo.Querier) Service {
	return &svc{q}
}

func (s *svc) CreateMovie(ctx context.Context, movie createMovieReq) (repo.CreateMovieRow, error) {
	releaseDate, err := time.Parse(time.RFC3339, movie.ReleaseDate)
	if err != nil {
		return repo.CreateMovieRow{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "release_date should be in iso time format",
		}
	}

	row, err := s.repo.CreateMovie(ctx, repo.CreateMovieParams{
		Name:        movie.Name,
		Description: movie.Description,
		Casts:       movie.Casts,
		TrailerUrl:  movie.Trailer,
		Languages:   movie.Languages,
		ReleaseDate: releaseDate,
		Director:    movie.Director,
		Status:      repo.ReleaseStatus(movie.Status),
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return repo.CreateMovieRow{}, err
		}

		body := make(map[string]string)
		code := http.StatusBadRequest
		if e.Code == "22P02" && strings.Contains(e.Message, "release_status") {
			body["release_date"] = "release_date must be either the values of 'RELEASED', 'UNRELEASED' or 'BLOCKED'"
		}

		return repo.CreateMovieRow{}, apiresponse.ApiError{StatusCode: code, Body: body}
	}

	return row, nil
}

func (s *svc) GetMovieById(ctx context.Context, id string) (repo.GetMovieByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return repo.GetMovieByIdRow{}, apiresponse.InvalidUUID()
	}

	row, err := s.repo.GetMovieById(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return repo.GetMovieByIdRow{}, apiresponse.ApiError{
				StatusCode: http.StatusNotFound,
				Body:       "No movies present for the given uuid",
			}
		}
		return repo.GetMovieByIdRow{}, err
	}

	return row, nil
}

func (s *svc) DeleteMovieById(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apiresponse.InvalidUUID()
	}

	numRows, err := s.repo.DeleteMovieById(ctx, uid)
	if err != nil {
		return err
	}

	if numRows == 0 {
		return apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid movie_id. The specified movie doesn't exists.",
		}
	}

	return nil
}

func (s *svc) UpdateMovieById(ctx context.Context, movie updateMovieReq, id string) (repo.UpdateMovieByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return repo.UpdateMovieByIdRow{}, apiresponse.InvalidUUID()
	}

	var releaseDate time.Time
	if movie.ReleaseDate != nil {
		releaseDate, err = time.Parse(time.RFC3339, *movie.ReleaseDate)
		if err != nil {
			return repo.UpdateMovieByIdRow{}, apiresponse.ApiError{
				StatusCode: http.StatusBadRequest,
				Body:       "release_date should be in iso time format",
			}
		}
	}

	var status repo.NullReleaseStatus
	if movie.Status != nil {
		status = repo.NullReleaseStatus{Valid: true, ReleaseStatus: repo.ReleaseStatus(*movie.Status)}
	}

	row, err := s.repo.UpdateMovieById(ctx, repo.UpdateMovieByIdParams{
		ID:          uid,
		Name:        db.ToNullString(movie.Name),
		Description: db.ToNullString(movie.Description),
		TrailerUrl:  db.ToNullString(movie.Trailer),
		Director:    db.ToNullString(movie.Director),
		ReleaseDate: pgtype.Timestamptz{Time: releaseDate, Valid: true},
		Status:      status,
		Casts:       movie.Casts,
		Languages:   movie.Languages,
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return repo.UpdateMovieByIdRow{}, err
		}

		body := make(map[string]string)
		code := http.StatusBadRequest
		if e.Code == "22P02" && strings.Contains(e.Message, "release_status") {
			body["release_date"] = "release_date must be either the values of 'RELEASED', 'UNRELEASED' or 'BLOCKED'"
		}

		return repo.UpdateMovieByIdRow{}, apiresponse.ApiError{StatusCode: code, Body: body}
	}

	return row, nil
}

func (s *svc) GetMoviesByName(ctx context.Context, query string, cityId uuid.UUID) ([]repo.GetMoviesByNameRow, error) {
	return s.repo.GetMoviesByName(ctx, query)
}

func (s *svc) GetUpcomingMovies(ctx context.Context, cityId uuid.UUID) ([]repo.GetUpcomingMoviesRow, error) {
	return s.repo.GetUpcomingMovies(ctx, cityId)
}

func (s *svc) GetAllMovies(ctx context.Context) ([]repo.GetAllMoviesRow, error) {
	return s.repo.GetAllMovies(ctx)
}
