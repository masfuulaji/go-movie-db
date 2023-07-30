package request

type CollectionCreateRequest struct {
    Name        string `json:"name" validate:"required"` 
    Description string `json:"description"`
    PosterPath  string `json:"poster_path"`
}
