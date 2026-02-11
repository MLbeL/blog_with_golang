package request

import (
	"log"
	"net/http"

	"github.com/MLbeL/blog_with_golang/pkg/response"
)

func Resp[essence any](w *http.ResponseWriter, r *http.Request) (essence, error) {
	body, err := Decode[essence](r.Body)
	if err != nil {
		log.Printf("error for decoding in req.go: %v\n", err)
		response.Json(err.Error(), *w, 402)
		return body, err
	}
	return body, nil
}
