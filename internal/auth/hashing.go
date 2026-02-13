package auth

import "golang.org/x/crypto/bcrypt"

const (
	Cost = 12
)

func CreateHashFromPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), Cost)
	if err != nil {
		return string(hash), err
	}
	return string(hash), err
}

func CompareHashToPassword(hash string, password string) error {
	answer := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return answer
}
