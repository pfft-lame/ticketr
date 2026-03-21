package theaters

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
	CreateTheater(ctx context.Context, theater createTheaterReq) (queries.CreateTheaterRow, error)
	DeleteTheaterById(ctx context.Context, id string) error
	UpdateTheaterById(ctx context.Context, id string, theater updateTheaterReq) (queries.UpdateTheatreByIdRow, error)
	GetTheaterById(ctx context.Context, id string) (queries.GetTheatersByIdRow, error)
	GetAllTheaters(ctx context.Context) ([]queries.GetAllTheatersRow, error)
	GetTheatersByCity(ctx context.Context, cityId string) ([]queries.GetTheatersByCityIdRow, error)
}

type svc struct {
	q queries.Querier
}

func NewService(q queries.Querier) Service {
	return &svc{q}
}

func (s *svc) CreateTheater(ctx context.Context, theater createTheaterReq) (queries.CreateTheaterRow, error) {
	cityId, _ := uuid.Parse(theater.CityId) // validated uuid

	row, err := s.q.CreateTheater(ctx, queries.CreateTheaterParams{
		Name:        theater.Name,
		Description: theater.Description,
		CityID:      cityId,
		Address:     theater.Address,
		Pincode:     theater.Pincode,
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return queries.CreateTheaterRow{}, err
		}

		code := http.StatusBadRequest
		body := make(map[string]string)
		if e.Code == db.ForeignKeyViolation && e.ConstraintName == "theaters_city_id_fkey" {
			body["city_id"] = "Invalid city_id. The specified city doesn't exists."
		}
		return queries.CreateTheaterRow{}, apiresponse.ApiError{
			StatusCode: code,
			Body:       body,
		}
	}

	return row, nil
}

func (s *svc) DeleteTheaterById(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apiresponse.InvalidUUID()
	}

	numRows, err := s.q.DeleteTheaterById(ctx, uid)
	if err != nil {
		return err
	}

	if numRows == 0 {
		return apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid theater_id. The specified theater doesn't exists.",
		}
	}

	return nil
}

func (s *svc) UpdateTheaterById(ctx context.Context, id string, theater updateTheaterReq) (queries.UpdateTheatreByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return queries.UpdateTheatreByIdRow{}, apiresponse.InvalidUUID()
	}

	var cityId pgtype.UUID
	if theater.CityId != nil {
		cityId.Scan(*theater.CityId) // validated uuid
	}

	row, err := s.q.UpdateTheatreById(ctx, queries.UpdateTheatreByIdParams{
		Name:        db.ToNullString(theater.Name),
		Description: db.ToNullString(theater.Description),
		CityID:      cityId,
		Address:     db.ToNullString(theater.Address),
		Pincode:     db.ToNullString(theater.Pincode),
		ID:          uid,
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return queries.UpdateTheatreByIdRow{}, err
		}

		code := http.StatusBadRequest
		body := make(map[string]string)
		if e.Code == db.ForeignKeyViolation && e.ColumnName == "theaters_city_id_fkey" {
			body["city_id"] = "Invalid city_id. The specified city doesn't exists."
		}

		return queries.UpdateTheatreByIdRow{}, apiresponse.ApiError{
			StatusCode: code,
			Body:       body,
		}
	}

	return row, nil
}

func (s *svc) GetTheaterById(ctx context.Context, id string) (queries.GetTheatersByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return queries.GetTheatersByIdRow{}, apiresponse.InvalidUUID()
	}

	return s.q.GetTheatersById(ctx, uid)
}

func (s *svc) GetAllTheaters(ctx context.Context) ([]queries.GetAllTheatersRow, error) {
	return s.q.GetAllTheaters(ctx)
}

func (s *svc) GetTheatersByCity(ctx context.Context, cityId string) ([]queries.GetTheatersByCityIdRow, error) {
	id, err := uuid.Parse(cityId)
	if err != nil {
		return nil, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid uuid for city",
		}
	}

	rows, err := s.q.GetTheatersByCityId(ctx, id)
	if err != nil {
		return nil, err
	}

	return rows, nil
}
