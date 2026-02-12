package db

import (
	"fmt"

	"github.com/MLbeL/blog_with_golang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDb(c *config.DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.DB_HOST, c.DB_USER, c.DB_PASSWORD, c.DB_NAME, c.DB_PORT, c.DB_SSLMODE,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
