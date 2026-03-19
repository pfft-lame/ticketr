package cities

type createCityReq struct {
	City  string `json:"city" validate:"required"`
	State string `json:"state" validate:"required"`
}
