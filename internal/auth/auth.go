package auth

import (
	"net/http"

	"github.com/MLbeL/blog_with_golang/db"
	"github.com/MLbeL/blog_with_golang/pkg/request"
	"github.com/MLbeL/blog_with_golang/pkg/response"
)

type AuthHandler struct {
	UserRepo *db.UserRepo
}

type AuthHandlerDeps struct {
	UserRepo *db.UserRepo
}

func NewHandlerFuncAuth(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{UserRepo: deps.UserRepo}
	router.HandleFunc("POST /auth/login", handler.LoginHandler())
	router.HandleFunc("POST /auth/register", handler.RegisterHandler())
}

func (au AuthHandler) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.Resp[Auth](&w, r)
		if err != nil {
			return
		}
		response.Json("Successful", w, 200)
	}
}

func (au AuthHandler) RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.Resp[Register](&w, r)
		if err != nil {
			return
		}
		// hash, err := CreateHashFromPassword(data.Password)
		// if err != nil {
		// 	response.Json("internal server error", w, 500)
		// 	return
		// }
		// user := &models.User{Name: data.Name, Email: data.Email, PasswordHash: hash}
		// err = au.UserRepo.CreateUser(user)
		response.Json("On your email was send letter with link...", w, 200)
	}
}
