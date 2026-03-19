package theater

type createTheaterReq struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	CityId      string `json:"city_id" validate:"required,uuid4"`
	Address     string `json:"address" validate:"required"`
	Pincode     string `json:"pincode" validate:"required,numeric,len=6"`
}

type updateTheaterReq struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
	CityId      *string `json:"city_id" validate:"omitempty,uuid4"`
	Address     *string `json:"address"`
	Pincode     *string `json:"pincode" validate:"omitempty,numeric,len=6"`
}
