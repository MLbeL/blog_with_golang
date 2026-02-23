package db

import "errors"

var (
	ErrPostNotFound   = errors.New("post not found")
	ErrUserIDNotFound = errors.New("user id not found")
)
