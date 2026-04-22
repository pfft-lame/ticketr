package theaters

import (
	"context"
	"fmt"
	"net/http"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/db"
	repo "ticketr/internal/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service interface {
	CreateTheater(ctx context.Context, theater createTheaterReq) (repo.CreateTheaterRow, error)
	DeleteTheaterById(ctx context.Context, id string) error
	UpdateTheaterById(ctx context.Context, id string, theater updateTheaterReq) (repo.UpdateTheatreByIdRow, error)
	GetTheaterById(ctx context.Context, id string) (repo.GetTheatersByIdRow, error)
	GetAllTheaters(ctx context.Context) ([]repo.GetAllTheatersRow, error)
	GetTheatersByCity(ctx context.Context, cityId uuid.UUID) ([]repo.GetTheatersByCityIdRow, error)
	GetUpcomingMoviesInTheater(ctx context.Context, id string, cityId uuid.UUID) ([]repo.GetUpcomingMoviesInTheaterRow, error)
}

type svc struct {
	q repo.Querier
}

func NewService(q repo.Querier) Service {
	return &svc{q}
}

func (s *svc) CreateTheater(ctx context.Context, theater createTheaterReq) (repo.CreateTheaterRow, error) {
	cityId, _ := uuid.Parse(theater.CityId) // validated uuid

	row, err := s.q.CreateTheater(ctx, repo.CreateTheaterParams{
		Name:        theater.Name,
		Description: theater.Description,
		CityID:      cityId,
		Address:     theater.Address,
		Pincode:     theater.Pincode,
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return repo.CreateTheaterRow{}, err
		}

		code := http.StatusBadRequest
		body := make(map[string]string)
		if e.Code == db.ForeignKeyViolation && e.ConstraintName == "theaters_city_id_fkey" {
			body["city_id"] = "Invalid city_id. The specified city doesn't exists."
		}
		return repo.CreateTheaterRow{}, apiresponse.ApiError{
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

func (s *svc) UpdateTheaterById(ctx context.Context, id string, theater updateTheaterReq) (repo.UpdateTheatreByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return repo.UpdateTheatreByIdRow{}, apiresponse.InvalidUUID()
	}

	var cityId pgtype.UUID
	if theater.CityId != nil {
		cityId.Scan(*theater.CityId) // validated uuid
	}

	row, err := s.q.UpdateTheatreById(ctx, repo.UpdateTheatreByIdParams{
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
			return repo.UpdateTheatreByIdRow{}, err
		}

		code := http.StatusBadRequest
		body := make(map[string]string)
		if e.Code == db.ForeignKeyViolation && e.ColumnName == "theaters_city_id_fkey" {
			body["city_id"] = "Invalid city_id. The specified city doesn't exists."
		}

		return repo.UpdateTheatreByIdRow{}, apiresponse.ApiError{
			StatusCode: code,
			Body:       body,
		}
	}

	return row, nil
}

func (s *svc) GetTheaterById(ctx context.Context, id string) (repo.GetTheatersByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return repo.GetTheatersByIdRow{}, apiresponse.InvalidUUID()
	}

	return s.q.GetTheatersById(ctx, uid)
}

func (s *svc) GetAllTheaters(ctx context.Context) ([]repo.GetAllTheatersRow, error) {
	return s.q.GetAllTheaters(ctx)
}

func (s *svc) GetTheatersByCity(ctx context.Context, cityId uuid.UUID) ([]repo.GetTheatersByCityIdRow, error) {
	rows, err := s.q.GetTheatersByCityId(ctx, cityId)
	if err != nil {
		return nil, err
	}

	return rows, nil
}

func (s *svc) GetUpcomingMoviesInTheater(ctx context.Context, id string, cityId uuid.UUID) ([]repo.GetUpcomingMoviesInTheaterRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		fmt.Println("here")
		return nil, apiresponse.InvalidUUID()
	}
	return s.q.GetUpcomingMoviesInTheater(ctx, repo.GetUpcomingMoviesInTheaterParams{
		ID:     uid,
		CityID: cityId,
	})
}
