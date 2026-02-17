package posts

import (
	"net/http"

	"github.com/MLbeL/blog_with_golang/db"
	"github.com/MLbeL/blog_with_golang/pkg/request"
)

type PostsHandler struct {
	PostRepo *db.PostRepo
}

type PostsHandlerDeps struct {
	PostRepo *db.PostRepo
}

func NewHandlerPosts(router *http.ServeMux, deps *PostsHandlerDeps) {
	handler := PostsHandler{PostRepo: deps.PostRepo}
	router.HandleFunc("GET /posts", handler.GET_Posts())
	router.HandleFunc("POST /posts", handler.POST_Posts())
	router.HandleFunc("GET /posts/{id}", handler.GETOnePost())
}

func (ph PostsHandler) GET_Posts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (ph PostsHandler) POST_Posts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.Resp[Post](&w, r)
		if err != nil {
			return
		}

	}
}

func (ph PostsHandler) GETOnePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
