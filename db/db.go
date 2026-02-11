package db

import (
	"fmt"

	"github.com/MLbeL/blog_with_golang/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func LoadDb(c *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Db.DB_HOST, c.Db.DB_USER, c.Db.DB_PASSWORD, c.Db.DB_NAME, c.Db.DB_PORT, c.Db.DB_SSLMODE,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
