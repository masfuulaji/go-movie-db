package request

type GenreCreateRequest struct {
    Name string `json:"name" validate:"required"`
    Description string `json:"description"`
}
