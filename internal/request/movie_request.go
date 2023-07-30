package request

type MovieCreateRequest struct {
	Title       string  `json:"title" validate:"required"`
	Overview    string  `json:"overview"`
	ReleaseDate string  `json:"release_date"`
	Runtime     int     `json:"runtime"`
	Budget      int     `json:"budget"`
	Revenue     int     `json:"revenue"`
	Popularity  float64 `json:"popularity"`
	PosterPath  string  `json:"poster_path"`
	Status      int     `json:"status"`
}
