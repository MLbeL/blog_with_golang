package main

import (
	"log"
	"net/http"

	"github.com/MLbeL/blog_with_golang/config"
	"github.com/MLbeL/blog_with_golang/db"
	"github.com/MLbeL/blog_with_golang/internal/auth"
)

func main() {
	router := http.NewServeMux()

	conf := config.GetConfig()
	DB, err := db.LoadDb(conf.Db)

	if err != nil {
		log.Fatalf("Failed to Load DB: %v", err)
	}
	log.Println("DB is loaded")

	secret := conf.JWTSecret
	if secret == "" {
		log.Fatal("JWT secret is empty")
	}

	userRepo := &db.UserRepo{DB: DB}
	auth.NewHandlerFuncAuth(router, &auth.AuthHandlerDeps{UserRepo: userRepo, Secret: secret})
	log.Println("Loaded a handler-funcs")
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Println("Server is starting on port 8080")
	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Server is broken with error: %v", err)
	}

}
