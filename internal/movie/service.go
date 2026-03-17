package movie

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/db"
	"ticketr/internal/db/queries"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	CreateMovie(context.Context, createMovieReq) (uuid.UUID, error)
	GetMovieById(context.Context, string) (queries.GetMovieByIdRow, error)
	DeleteMovieById(context.Context, string) error
	UpdateMovieById(context.Context, updateMovieReq, string) (queries.UpdateMovieByIdRow, error)
}

type svc struct {
	queries queries.Querier
}

func NewService(q queries.Querier) Service {
	return &svc{q}
}

func (s *svc) CreateMovie(ctx context.Context, movie createMovieReq) (uuid.UUID, error) {
	releaseDate, err := time.Parse(time.RFC3339, movie.ReleaseDate)
	if err != nil {
		return uuid.UUID{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "release_date should be in iso time format",
		}
	}

	id, err := s.queries.CreateMovie(ctx, queries.CreateMovieParams{
		Name:        movie.Name,
		Description: movie.Description,
		Casts:       movie.Casts,
		TrailerUrl:  movie.Trailer,
		Languages:   movie.Languages,
		ReleaseDate: releaseDate,
		Director:    movie.Director,
		Status:      queries.ReleaseStatus(movie.Status),
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return uuid.UUID{}, err
		}

		body := make(map[string]string)
		code := http.StatusBadRequest
		if e.Code == "22P02" && strings.Contains(e.Message, "release_status") {
			body["release_date"] = "release_date must be either the values of 'RELEASED', 'UNRELEASED' or 'BLOCKED'"
		}

		return uuid.UUID{}, apiresponse.ApiError{StatusCode: code, Body: body}
	}

	return id, nil
}

func (s *svc) GetMovieById(ctx context.Context, id string) (queries.GetMovieByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return queries.GetMovieByIdRow{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid UUID",
		}
	}

	row, err := s.queries.GetMovieById(ctx, uid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return queries.GetMovieByIdRow{}, apiresponse.ApiError{
				StatusCode: http.StatusNotFound,
				Body:       "No movies present for the given uuid",
			}
		}
		return queries.GetMovieByIdRow{}, err
	}

	return row, nil
}

func (s *svc) DeleteMovieById(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid UUID",
		}
	}

	err = s.queries.DeleteMovieById(ctx, uid)
	if err != nil {
		return err
	}

	return nil
}

func (s *svc) UpdateMovieById(ctx context.Context, movie updateMovieReq, id string) (queries.UpdateMovieByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return queries.UpdateMovieByIdRow{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid UUID",
		}
	}

	var releaseDate time.Time
	if movie.ReleaseDate != nil {
		releaseDate, err = time.Parse(time.RFC3339, *movie.ReleaseDate)
		if err != nil {
			return queries.UpdateMovieByIdRow{}, apiresponse.ApiError{
				StatusCode: http.StatusBadRequest,
				Body:       "release_date should be in iso time format",
			}
		}
	}

	var status queries.NullReleaseStatus
	if movie.Status != nil {
		status = queries.NullReleaseStatus{Valid: true, ReleaseStatus: queries.ReleaseStatus(*movie.Status)}
	}

	row, err := s.queries.UpdateMovieById(ctx, queries.UpdateMovieByIdParams{
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
			return queries.UpdateMovieByIdRow{}, err
		}

		body := make(map[string]string)
		code := http.StatusBadRequest
		if e.Code == "22P02" && strings.Contains(e.Message, "release_status") {
			body["release_date"] = "release_date must be either the values of 'RELEASED', 'UNRELEASED' or 'BLOCKED'"
		}

		return queries.UpdateMovieByIdRow{}, apiresponse.ApiError{StatusCode: code, Body: body}
	}

	return row, nil
}
