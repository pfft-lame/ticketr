package cities

import (
	"context"
	"net/http"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/db"
	"ticketr/internal/db/queries"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type Service interface {
	CreateCity(ctx context.Context, city createCityReq) (queries.CreateCityRow, error)
	DeleteCity(ctx context.Context, id string) error
	GetCityById(ctx context.Context, id string) (queries.GetCityByIdRow, error)
	GetAllCities(ctx context.Context) ([]queries.GetAllCitiesRow, error)
}

type svc struct {
	q queries.Querier
}

func NewService(q queries.Querier) Service {
	return &svc{q}
}

func (s *svc) CreateCity(ctx context.Context, city createCityReq) (queries.CreateCityRow, error) {
	row, err := s.q.CreateCity(ctx, queries.CreateCityParams{
		City:  city.City,
		State: city.State,
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return queries.CreateCityRow{}, err
		}

		code := http.StatusBadRequest
		body := make(map[string]string)

		if e.Code == db.UniqueConstraintViolation {
			body["city"] = "city-state pair already exists!"
		}

		return queries.CreateCityRow{}, apiresponse.ApiError{
			StatusCode: code,
			Body:       body,
		}
	}

	return row, nil
}

func (s *svc) DeleteCity(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apiresponse.InvalidUUID()
	}

	numRows, err := s.q.DeleteCityById(ctx, uid)
	if err != nil {
		return err
	}

	if numRows == 0 {
		return apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "No city with the given city_id was found.",
		}
	}

	return nil
}

func (s *svc) GetCityById(ctx context.Context, id string) (queries.GetCityByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return queries.GetCityByIdRow{}, apiresponse.InvalidUUID()
	}

	row, err := s.q.GetCityById(ctx, uid)
	if err != nil {
		return queries.GetCityByIdRow{}, err
	}

	return row, nil
}

func (s *svc) GetAllCities(ctx context.Context) ([]queries.GetAllCitiesRow, error) {
	return s.q.GetAllCities(ctx)
}
