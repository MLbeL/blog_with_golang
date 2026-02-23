package auth

import "time"

type Auth struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,alphanum"`
	Name     string `json:"name" validate:"min=2"`
}

type RefreshToken struct {
	Refreshtoken string
}

const RefreshTokenTTL = 30 * 24 * time.Hour
