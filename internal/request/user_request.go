package request

type UserCreateRequest struct {
    Name string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required"`
}

type UserUpdateRequest struct {
    Name string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required"`
    Password string `json:"password"`
}

type UserLoginRequest struct {
    Email string `json:"email" validate:"required"`
    Password string `json:"password" validate:"required"`
}
