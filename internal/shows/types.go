package shows

import (
	"time"

	"github.com/google/uuid"
)

type createShowReq struct {
	MovieId   uuid.UUID `json:"movie_id" validate:"required,uuid4"`
	ScreenId  uuid.UUID `json:"screen_id" validate:"required,uuid4"`
	StartTime time.Time `json:"start_time" validate:"required"`
	EndTime   time.Time `json:"end_time" validate:"required"`
}

type updateShowReq struct {
	MovieId   *uuid.UUID
	ScreenId  *uuid.UUID
	StartTime *time.Time
	EndTime   *time.Time
}
