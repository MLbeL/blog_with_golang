package posts

import (
	"time"

	"github.com/MLbeL/blog_with_golang/db/models"
)

type Post struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

type ShowPost struct {
	IDAuthor  uint      `json:"authorid"`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"createdat"`
}

type UpdatePost struct {
	Title *string `json:"title,omitempty"`
	Text  *string `json:"text,omitempty"`
}

func ChangeTypeToShowPost(posts []models.Post) []ShowPost {
	newType := make([]ShowPost, len(posts))

	for i, v := range posts {
		newType[i] = ShowPost{
			IDAuthor:  v.AuthorID,
			Title:     v.Title,
			Text:      v.Text,
			CreatedAt: v.CreatedAt,
		}
	}
	return newType
}

func ChangeOnePostToShow(post models.Post) ShowPost {
	return ShowPost{
		IDAuthor:  post.AuthorID,
		Title:     post.Title,
		Text:      post.Text,
		CreatedAt: post.CreatedAt,
	}
}
