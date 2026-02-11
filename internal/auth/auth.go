package auth

import (
	"net/http"

	"github.com/MLbeL/blog_with_golang/pkg/request"
	"github.com/MLbeL/blog_with_golang/pkg/response"
)

type AuthHandler struct {
}

type AuthHandlerDeps struct {
}

func NewHandlerFunc(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{}
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

	}
}
