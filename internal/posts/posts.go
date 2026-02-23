package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/MLbeL/blog_with_golang/db"
	"github.com/MLbeL/blog_with_golang/db/models"
	"github.com/MLbeL/blog_with_golang/pkg/request"
	"github.com/MLbeL/blog_with_golang/pkg/response"
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
	router.HandleFunc("PATCH /posts/{id}", handler.UpdatePost())
	router.HandleFunc("DELETE /posts/{id}", handler.DeletePost())
}

func (ph PostsHandler) GET_Posts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := ph.PostRepo.ShowAllPosts()
		if err != nil {
			response.Json("internal server error", w, 500)
			return
		}
		arr_posts := ChangeTypeToShowPost(posts) // change type for json-answer
		response.Json(arr_posts, w, 200)
	}
}

func (ph PostsHandler) POST_Posts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.Resp[Post](&w, r)
		if err != nil {
			return
		}
		model_post := models.Post{
			AuthorID: 1, // TODO: replace with real user id from auth
			Title:    body.Title,
			Text:     body.Text,
		}
		err = ph.PostRepo.CreatePost(&model_post)
		if err != nil {
			response.Json("Internal server error", w, 500)
			return
		}
		response.Json("Created new post", w, 201)
	}
}

func (ph PostsHandler) GETOnePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			response.Json("internal server error", w, 500)
			return
		}
		post, err := ph.PostRepo.FindPostByID(uint(id))
		if err != nil {
			if errors.Is(err, db.ErrPostNotFound) {
				response.Json("post not found", w, 404)
				return
			} else {
				response.Json("internal server error", w, 500)
				return
			}
		}
		postJson := ChangeOnePostToShow(post)
		response.Json(postJson, w, 200)
	}
}

func (ph PostsHandler) UpdatePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			response.Json("internal server error", w, 500)
			return
		}
		body, err := request.Resp[UpdatePost](&w, r)
		if err != nil {
			return
		}
		fields := make(map[string]interface{})
		if body.Title != nil {
			fields["title"] = *body.Title
		}
		if body.Text != nil {
			fields["text"] = *body.Text
		}
		err = ph.PostRepo.UpdatePost(uint(id), fields)
		if err != nil {
			if errors.Is(err, db.ErrPostNotFound) {
				response.Json("post not found", w, 404)
				return
			} else {
				response.Json("internal server error", w, 500)
				return
			}
		}
		response.Json("Succesful updated", w, 200)
	}
}

func (ph PostsHandler) DeletePost() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 64)
		if err != nil {
			response.Json("internal server error", w, 500)
			return
		}
		err = ph.PostRepo.DeletePost(uint(id))
		if err != nil {
			if errors.Is(err, db.ErrPostNotFound) {
				response.Json("post with this id not found", w, 404)
				return
			} else {
				response.Json("internal server error", w, 500)
				return
			}
		}

		response.Json("Succesful delete post", w, 200)

	}
}
