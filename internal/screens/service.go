package screens

import (
	"context"
	"net/http"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/db"
	"ticketr/internal/db/queries"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	CreateScreen(ctx context.Context, screen createScreenReq) (queries.CreateScreenRow, error)
	UpdateScreenId(ctx context.Context, id string, screen updateScreenReq) (queries.UpdateScreenByIdRow, error)
	DeleteScreenById(ctx context.Context, id string) error
	GetScreenById(ctx context.Context, id string) (queries.GetScreenByIdRow, error)
	GetScreenByTheaterId(ctx context.Context, theaterID string) ([]queries.GetAllScreensByTheaterIdRow, error)
}

type svc struct {
	q queries.Querier
}

func NewService(q queries.Querier) Service {
	return &svc{q}
}

func (s *svc) CreateScreen(ctx context.Context, screen createScreenReq) (queries.CreateScreenRow, error) {
	theaterId, _ := uuid.Parse(screen.TheaterId) // validated uuid

	row, err := s.q.CreateScreen(ctx, queries.CreateScreenParams{
		Name:       screen.Name,
		TotalSeats: int32(screen.TotalSeats),
		TheaterID:  theaterId,
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return queries.CreateScreenRow{}, err
		}

		body := make(map[string]string)
		if e.Code == db.UniqueConstraintViolation && e.ConstraintName == "unique_theater_id_name" {
			body["name"] = "Screen with given name already exists."
		}
		if e.Code == db.ForeignKeyViolation && e.ConstraintName == "screens_theater_id_fkey" {
			body["theater_id"] = "Invalid theater_id. The specified theater doesn't exists."
		}

		return queries.CreateScreenRow{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       body,
		}
	}

	return row, nil
}

func (s *svc) UpdateScreenId(ctx context.Context, id string, screen updateScreenReq) (queries.UpdateScreenByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return queries.UpdateScreenByIdRow{}, apiresponse.InvalidUUID()
	}

	var theaterId pgtype.UUID
	if screen.TheaterId != nil {
		theaterId.Scan(*screen.TheaterId) // validated uuid
	}

	row, err := s.q.UpdateScreenById(ctx, queries.UpdateScreenByIdParams{
		Name:       db.ToNullString(screen.Name),
		TotalSeats: db.ToNullInt32(screen.TotalSeats),
		TheaterID:  theaterId,
		ID:         uid,
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return queries.UpdateScreenByIdRow{}, err
		}

		body := make(map[string]string)
		if e.Code == db.UniqueConstraintViolation && e.ConstraintName == "unique_theater_id_name" {
			body["name"] = "Screen with given name already exists."
		}
		if e.Code == db.ForeignKeyViolation && e.ConstraintName == "screens_theater_id_fkey" {
			body["theater_id"] = "Invalid theater_id. The specified theater doesn't exists."
		}

		return queries.UpdateScreenByIdRow{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       body,
		}
	}

	return row, nil
}

func (s *svc) DeleteScreenById(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apiresponse.InvalidUUID()
	}

	numRows, err := s.q.DeleteScreenByID(ctx, uid)
	if err != nil {
		return err
	}

	if numRows == 0 {
		return apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid screen_id. The specified screen doesn't exists.",
		}
	}

	return nil
}

func (s *svc) GetScreenById(ctx context.Context, id string) (queries.GetScreenByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return queries.GetScreenByIdRow{}, apiresponse.InvalidUUID()
	}

	return s.q.GetScreenById(ctx, uid)
}

func (s *svc) GetScreenByTheaterId(ctx context.Context, theaterID string) ([]queries.GetAllScreensByTheaterIdRow, error) {
	uid, err := uuid.Parse(theaterID)
	if err != nil {
		return nil, apiresponse.InvalidUUID()
	}

	return s.q.GetAllScreensByTheaterId(ctx, uid)
}
