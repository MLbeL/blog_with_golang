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

func (ur *UserRepo) GetHashPasswordByEmailAndUserID(email string) (string, uint, error) {
	var user models.User
	err := ur.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", 0, err
		}
		return "", 0, err
	}
	return user.PasswordHash, user.ID, nil
}

// func (ur *UserRepo) GetUserIDByEmail(email string) (uint, error) {
// 	var user models.User
// 	err := ur.DB.First(&user, "email=?", email)
// 	if err != nil {
// 		return 0, err.Error
// 	}
// 	return user.BaseModel.ID, nil
// }

func (ur *UserRepo) SaveRefreshTokenFromDB(rt *models.RefreshToken) error {
	return ur.DB.Create(rt).Error
}
