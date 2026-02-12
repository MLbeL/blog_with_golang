package auth

type Auth struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email" validate:"required, email"`
	Password string `json:"password" validate:"required, min=8, alphanum"`
	Name     string `json:"name" validate:"min=2"`
}
