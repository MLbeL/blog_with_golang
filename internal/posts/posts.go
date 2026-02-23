package posts

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/MLbeL/blog_with_golang/db"
	"github.com/MLbeL/blog_with_golang/db/models"
	"github.com/MLbeL/blog_with_golang/pkg/middleware"
	"github.com/MLbeL/blog_with_golang/pkg/request"
	"github.com/MLbeL/blog_with_golang/pkg/response"
)

type PostsHandler struct {
	PostRepo *db.PostRepo
	UserRepo *db.UserRepo
}

type PostsHandlerDeps struct {
	PostRepo *db.PostRepo
	UserRepo *db.UserRepo
}

func NewHandlerPosts(router *http.ServeMux, deps *PostsHandlerDeps, secret string) {
	handler := PostsHandler{PostRepo: deps.PostRepo, UserRepo: deps.UserRepo}
	router.HandleFunc("GET /posts", handler.GET_Posts())
	router.Handle("POST /posts", middleware.CheckAuth(handler.POST_Posts(), secret))
	router.HandleFunc("GET /posts/{id}", handler.GETOnePost())
	router.Handle("PATCH /posts/{id}", middleware.CheckAuth(handler.UpdatePost(), secret))
	router.Handle("DELETE /posts/{id}", middleware.CheckAuth(handler.DeletePost(), secret))
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
		email := r.Context().Value(middleware.CtxEmailKey).(string)
		log.Println(email)
		_, userID, err := ph.UserRepo.GetHashPasswordByEmailAndUserID(email)
		log.Println(userID)
		if err != nil {
			if errors.Is(err, db.ErrUserIDNotFound) {
				response.Json("Unauthorized", w, 401)
				return
			}
			response.Json("Internal server error", w, 500)
			return
		}

		model_post := models.Post{
			AuthorID: userID,
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

		// Check if user is author of post
		email := r.Context().Value(middleware.CtxEmailKey).(string)
		_, userID, err := ph.UserRepo.GetHashPasswordByEmailAndUserID(email)
		if err != nil {
			if errors.Is(err, db.ErrUserIDNotFound) {
				response.Json("Unauthorized", w, 401)
				return
			}
			response.Json("Internal server error", w, 500)
			return
		}
		post, err := ph.PostRepo.FindPostByID(uint(id))
		if err != nil {
			response.Json("internal server error", w, 500)
			return
		}
		if userID != post.AuthorID {
			response.Json("No permission", w, 403)
			return
		}
		// end of check user is author

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

		// Check if user is author of post
		email := r.Context().Value(middleware.CtxEmailKey).(string)
		_, userID, err := ph.UserRepo.GetHashPasswordByEmailAndUserID(email)
		if err != nil {
			if errors.Is(err, db.ErrUserIDNotFound) {
				response.Json("Unauthorized", w, 401)
				return
			}
			response.Json("Internal server error", w, 500)
			return
		}
		post, err := ph.PostRepo.FindPostByID(uint(id))
		if err != nil {
			if errors.Is(err, db.ErrPostNotFound) {
				response.Json("Post not found", w, 404)
				return
			}
			response.Json("Internal server error", w, 500)
			return
		}
		if userID != post.AuthorID {
			response.Json("No permission", w, 403)
			return
		}
		// end of check user is author

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
