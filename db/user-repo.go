package db

import (
	"github.com/MLbeL/blog_with_golang/db/models"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB *gorm.DB
}

func (ur *UserRepo) CreateUser(u *models.User) error {
	return ur.DB.Create(u).Error
}
