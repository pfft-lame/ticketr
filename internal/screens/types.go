package screens

type createScreenReq struct {
	Name       string `json:"name" validate:"required"`
	TheaterId  string `json:"theater_id" validate:"required,uuid4"`
	TotalSeats int    `json:"total_seats" validate:"required"`
}

type updateScreenReq struct {
	Name       *string `json:"name"`
	TheaterId  *string `json:"theater_id" validate:"omitempty,uuid4"`
	TotalSeats *int    `json:"total_seats"`
}
