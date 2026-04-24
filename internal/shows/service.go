package shows

import (
	"context"
	"net/http"

	apiresponse "ticketr/internal/api_response"
	"ticketr/internal/db"
	repo "ticketr/internal/repository"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
)

type Service interface {
	CreateShow(ctx context.Context, show createShowReq) (repo.CreateShowRow, error)
	UpdateShowById(ctx context.Context, show updateShowReq, id string) (repo.UpdateShowByIdRow, error)
	DeleteShowById(ctx context.Context, id string) error
	GetShowInfoById(ctx context.Context, id string) (repo.GetShowInfoByIdRow, error)
	GetShowsByMovieId(ctx context.Context, movieId string, cityId uuid.UUID) ([]repo.GetShowsByMovieIdRow, error)
	GetShowsByTheaterId(ctx context.Context, theaterId string) ([]repo.GetShowsByTheaterIdRow, error)
	GetShowsByCityId(ctx context.Context, cityId uuid.UUID) ([]repo.GetShowsByCityIdRow, error)
}

type svc struct {
	q repo.Querier
}

func NewService(q repo.Querier) Service {
	return &svc{q}
}

func (s *svc) CreateShow(ctx context.Context, show createShowReq) (repo.CreateShowRow, error) {
	res, err := s.q.CreateShow(ctx, repo.CreateShowParams{
		ScreenID:  show.ScreenId,
		MovieID:   show.MovieId,
		StartTime: show.StartTime,
		EndTime:   show.EndTime,
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return repo.CreateShowRow{}, err
		}

		body := make(map[string]string)
		if e.Code == db.ForeignKeyViolation && e.ConstraintName == "shows_movie_id_fkey" {
			body["movie_id"] = "Invalid movie_id. The specified movie doesn't exists."
		}
		if e.Code == db.ForeignKeyViolation && e.ConstraintName == "shows_screen_id_fkey" {
			body["screen_id"] = "Invalid screen_id. The specified screen doesn't exists."
		}
		if e.Code == db.CheckViolation && e.ConstraintName == "valid_time" {
			body["start_time"] = "start_time should always be lesser than end_time"
		}

		return repo.CreateShowRow{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       body,
		}
	}

	return res, nil
}

func (s *svc) UpdateShowById(ctx context.Context, show updateShowReq, id string) (repo.UpdateShowByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return repo.UpdateShowByIdRow{}, apiresponse.InvalidUUID()
	}

	res, err := s.q.UpdateShowById(ctx, repo.UpdateShowByIdParams{
		ID:        uid,
		MovieID:   db.ToPgUUID(show.MovieId),
		ScreenID:  db.ToPgUUID(show.ScreenId),
		StartTime: db.ToPgTimeTz(show.StartTime),
		EndTime:   db.ToPgTimeTz(show.EndTime),
	})
	if err != nil {
		e, ok := err.(*pgconn.PgError)
		if !ok {
			return repo.UpdateShowByIdRow{}, err
		}

		body := make(map[string]string)
		if e.Code == db.ForeignKeyViolation && e.ConstraintName == "shows_movie_id_fkey" {
			body["movie_id"] = "Invalid movie_id. The specified movie doesn't exists."
		}
		if e.Code == db.ForeignKeyViolation && e.ConstraintName == "shows_screen_id_fkey" {
			body["screen_id"] = "Invalid screen_id. The specified screen doesn't exists."
		}
		if e.Code == db.CheckViolation && e.ConstraintName == "valid_time" {
			body["start_time"] = "start_time should always be lesser than end_time"
		}

		return repo.UpdateShowByIdRow{}, apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       body,
		}
	}

	return res, nil
}

func (s *svc) DeleteShowById(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return apiresponse.InvalidUUID()
	}

	numRows, err := s.q.DeleteShowById(ctx, uid)
	if err != nil {
		return err
	}

	if numRows == 0 {
		return apiresponse.ApiError{
			StatusCode: http.StatusBadRequest,
			Body:       "Invalid show_id. The specified show doesn't exists.",
		}
	}

	return nil
}

func (s *svc) GetShowInfoById(ctx context.Context, id string) (repo.GetShowInfoByIdRow, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return repo.GetShowInfoByIdRow{}, apiresponse.InvalidUUID()
	}

	res, err := s.q.GetShowInfoById(ctx, uid)
	if err != nil {
		return repo.GetShowInfoByIdRow{}, err
	}

	return res, nil
}

func (s *svc) GetShowsByMovieId(ctx context.Context, movieId string, cityId uuid.UUID) ([]repo.GetShowsByMovieIdRow, error) {
	mId, err := uuid.Parse(movieId)
	if err != nil {
		return nil, apiresponse.InvalidUUID()
	}

	res, err := s.q.GetShowsByMovieId(ctx, repo.GetShowsByMovieIdParams{
		CityID:  cityId,
		MovieID: mId,
	})
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *svc) GetShowsByTheaterId(ctx context.Context, theaterId string) ([]repo.GetShowsByTheaterIdRow, error) {
	tId, err := uuid.Parse(theaterId)
	if err != nil {
		return nil, apiresponse.InvalidUUID()
	}

	res, err := s.q.GetShowsByTheaterId(ctx, tId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *svc) GetShowsByCityId(ctx context.Context, cityId uuid.UUID) ([]repo.GetShowsByCityIdRow, error) {
	return s.q.GetShowsByCityId(ctx, cityId)
}
