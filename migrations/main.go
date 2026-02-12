package main

import (
	"log"

	"github.com/MLbeL/blog_with_golang/config"
	"github.com/MLbeL/blog_with_golang/db"
	"github.com/MLbeL/blog_with_golang/models"
)

func main() {
	confDB := config.GetConfig().Db
	DB, err := db.LoadDb(confDB)
	if err != nil {
		log.Fatalf("failed to load database: %v", err)
	}
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}
