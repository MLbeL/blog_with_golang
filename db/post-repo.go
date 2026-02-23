package db

import (
	"errors"

	"github.com/MLbeL/blog_with_golang/db/models"
	"gorm.io/gorm"
)

type PostRepo struct {
	DB *gorm.DB
}

func (pr *PostRepo) CreatePost(post *models.Post) error {
	return pr.DB.Create(post).Error
}

func (pr *PostRepo) ShowAllPosts() ([]models.Post, error) {
	var posts []models.Post
	result := pr.DB.Find(&posts)
	if result.Error != nil {
		return nil, result.Error
	}
	return posts, nil
}

func (pr *PostRepo) FindPostByID(id uint) (models.Post, error) {
	var post models.Post

	result := pr.DB.First(&post, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return models.Post{}, ErrPostNotFound
	}
	if result.Error != nil {
		return models.Post{}, result.Error
	}
	return post, nil
}

func (pr *PostRepo) UpdatePost(id uint, fields map[string]interface{}) error {
	result := pr.DB.Model(&models.Post{}).Where("id = ?", id).Updates(fields)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return ErrPostNotFound
	}
	return nil
}

func (pr *PostRepo) DeletePost(id uint) error {
	result := pr.DB.Where("id = ?", id).Delete(&models.Post{})
	if result.RowsAffected == 0 {
		return ErrPostNotFound
	}
	if result.Error != nil {
		return result.Error
	}
	return nil
}
