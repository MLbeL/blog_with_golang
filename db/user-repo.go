package db

import (
	"errors"

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
			return "", 0, ErrUserIDNotFound
		}
		return "", 0, err
	}
	return user.PasswordHash, user.ID, nil
}

func (ur *UserRepo) GetEmailByUserID(id uint) (string, error) {
	var user models.User
	err := ur.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", ErrUserIDNotFound
		}
		return "", err
	}
	return user.Email, nil
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

func (ur *UserRepo) DeleteRefreshTokensForUser(userID uint) error {
	return ur.DB.Where("user_id = ?", userID).Delete(&models.RefreshToken{}).Error
}

func (ur *UserRepo) FindUserIDByRefreshToken(rt string) (uint, error) {
	var user models.RefreshToken

	err := ur.DB.Where("token = ?", rt).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return 0, ErrUserIDNotFound
		} else {
			return 0, err
		}
	}
	return user.UserID, nil
}
