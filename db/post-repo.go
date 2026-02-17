package db

import "gorm.io/gorm"

type PostRepo struct {
	db *gorm.DB
}
