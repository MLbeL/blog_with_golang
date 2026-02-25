package main

import (
	"log"
	"time"

	"github.com/MLbeL/blog_with_golang/config"
	"github.com/MLbeL/blog_with_golang/db"
	"github.com/MLbeL/blog_with_golang/db/models"
	"gorm.io/gorm"
)

func main() {
	confDB := config.GetConfig().Db

	var (
		DB  *gorm.DB
		err error
	)

	for i := 0; i < 10; i++ {
		DB, err = db.LoadDb(confDB)
		if err == nil {
			break
		}
		log.Printf("failed to load database (attempt %d/10): %v", i+1, err)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("failed to load database after retries: %v", err)
	}

	if err := DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.RefreshToken{}); err != nil {
		log.Fatalf("failed to run migrations: %v", err)
	}
}
