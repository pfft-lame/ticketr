package movies

type createMovieReq struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Casts       []string `json:"casts" validate:"required,min=1,dive,required,min=2"`
	Trailer     string   `json:"trailer" validate:"required"`
	Languages   []string `json:"languages" validate:"required,min=1,dive,required,min=2"`
	Director    string   `json:"director" validate:"required"`
	ReleaseDate string   `json:"release_date" validate:"required"`
	Status      string   `json:"status" validate:"required,oneof=UNRELEASED RELEASED BLOCKED"`
}

/*
validate:"required,min=1,dive,required,min=2"
ORDER MATTERS. here dive means the value should be a iterable. so if we put min=1 after dive it checks if all the values withing list is atleast 1 char long
above validate tag says that the field is required and have minimum size of 1 and dive type again whose elements are required (so can't pass empty list)
and the size of each elements must atleast be 2
*/

type updateMovieReq struct {
	Name        *string  `json:"name"`
	Description *string  `json:"description"`
	Casts       []string `json:"casts" validate:"omitempty,min=1,dive,required,min=2"`
	Trailer     *string  `json:"trailer"`
	Languages   []string `json:"languages" validate:"omitempty,min=1,dive,required,min=2"`
	Director    *string  `json:"director"`
	ReleaseDate *string  `json:"release_date"`
	Status      *string  `json:"status" validate:"omitempty,oneof=UNRELEASED RELEASED BLOCKED"`
}
