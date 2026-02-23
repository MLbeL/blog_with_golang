package auth

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/MLbeL/blog_with_golang/db"
	"github.com/MLbeL/blog_with_golang/db/models"
	"github.com/MLbeL/blog_with_golang/pkg/request"
	"github.com/MLbeL/blog_with_golang/pkg/response"
	"gorm.io/gorm"
)

type AuthHandler struct {
	UserRepo *db.UserRepo
	Secret   string
}

type AuthHandlerDeps struct {
	UserRepo *db.UserRepo
	Secret   string
}

func NewHandlerFuncAuth(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{UserRepo: deps.UserRepo, Secret: deps.Secret}
	router.HandleFunc("POST /auth/login", handler.LoginHandler())
	router.HandleFunc("POST /auth/register", handler.RegisterHandler())
	router.HandleFunc("POST /auth/refresh", handler.NewAccessTokenByRefreshHandler())
}

func (au AuthHandler) LoginHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := request.Resp[Auth](&w, r)
		if err != nil {
			return
		}

		hash_db, userID, err_db := au.UserRepo.GetHashPasswordByEmailAndUserID(data.Email)
		if err_db != nil {
			if errors.Is(err_db, gorm.ErrRecordNotFound) {
				response.Json("Incorrect email or password", w, 401)
				return
			}
			response.Json("internal server error", w, 500)
			log.Println("DB error: ", err_db)
			return
		}

		isCorrectPassword := CompareHashToPassword(hash_db, data.Password)
		if isCorrectPassword != nil {
			response.Json("Incorrect email or password", w, 401)
			return
		}

		Accesstoken, err := GenerateAccessToken(userID, au.Secret, data.Email)
		if err != nil {
			response.Json("internal server error", w, 500)
			log.Println("GenerateAccessToken error: ", err)
			return
		}

		errDeleteRefresh := au.UserRepo.DeleteRefreshTokensForUser(userID)
		if errDeleteRefresh != nil {
			response.Json("internal server error", w, 500)
			log.Printf("Error for deleting refresh tokens for user with ID %v\n", userID)
			return
		}

		Refreshtoken, err := GenerateRefreshToken()
		if err != nil {
			response.Json("internal server error", w, 500)
			log.Println("GenerateRefreshToken error: ", err)
			return
		}
		SaveRefreshToken := &models.RefreshToken{
			UserID:    userID,
			Token:     Refreshtoken,
			ExpiresAt: time.Now().Add(RefreshTokenTTL),
		}
		err = au.UserRepo.SaveRefreshTokenFromDB(SaveRefreshToken)
		if err != nil {
			response.Json("internal server error", w, 500)
			log.Println("DB error: ", err)
			return
		}

		response.Json(map[string]string{
			"access_token":  Accesstoken,
			"refresh_token": Refreshtoken,
		}, w, 200)
	}
}

func (au AuthHandler) RegisterHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := request.Resp[Register](&w, r)
		if err != nil {
			return
		}
		hash, err := CreateHashFromPassword(data.Password)
		if err != nil {
			response.Json("internal server error", w, 500)
			return
		}
		user := &models.User{Name: data.Name, Email: data.Email, PasswordHash: hash}
		err = au.UserRepo.CreateUser(user)
		response.Json("Successful register!", w, 200)
	}
}

func (au AuthHandler) NewAccessTokenByRefreshHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.Resp[RefreshToken](&w, r)
		if err != nil {
			return
		}
		userID, err := au.UserRepo.FindUserIDByRefreshToken(body.Refreshtoken)
		if err != nil {
			if errors.Is(err, db.ErrUserIDNotFound) {
				response.Json("Invalid token(userID not found in db)", w, 401)
				return
			} else {
				response.Json("internal server error", w, 500)
				return
			}
		}
		email, err := au.UserRepo.GetEmailByUserID(userID)
		if err != nil {
			if errors.Is(err, db.ErrUserIDNotFound) {
				response.Json("Invalid token(userID not found in db)", w, 401)
				return
			} else {
				response.Json("internal server error", w, 500)
				return
			}
		}
		Accesstoken, err := GenerateAccessToken(userID, au.Secret, email)
		if err != nil {
			response.Json("internal server error", w, 500)
			log.Println("GenerateAccessToken error: ", err)
			return
		}

		errDeleteRefresh := au.UserRepo.DeleteRefreshTokensForUser(userID)
		if errDeleteRefresh != nil {
			response.Json("internal server error", w, 500)
			log.Printf("Error for deleting refresh tokens for user with ID %v\n", userID)
			return
		}

		Refreshtoken, err := GenerateRefreshToken()
		if err != nil {
			response.Json("internal server error", w, 500)
			log.Println("GenerateRefreshToken error: ", err)
			return
		}
		SaveRefreshToken := &models.RefreshToken{
			UserID:    userID,
			Token:     Refreshtoken,
			ExpiresAt: time.Now().Add(RefreshTokenTTL),
		}
		err = au.UserRepo.SaveRefreshTokenFromDB(SaveRefreshToken)
		if err != nil {
			response.Json("internal server error", w, 500)
			log.Println("DB error: ", err)
			return
		}

		response.Json(map[string]string{
			"access_token":  Accesstoken,
			"refresh_token": Refreshtoken,
		}, w, 200)
	}
}
