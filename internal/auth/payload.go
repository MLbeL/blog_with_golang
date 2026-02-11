package auth

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Register struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}
