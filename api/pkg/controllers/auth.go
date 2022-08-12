package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	models "github.com/Djancyp/go-rest/pkg/models"
	"github.com/Djancyp/go-rest/pkg/utils"
	"github.com/golang-jwt/jwt/v4"
)

// Create the JWT key used to create the signature
// TODO: move my_secret_key to .env
var jwtKey = []byte("my_secret_key")

// TODO: get value from .env
var expirationTime = time.Now().Add(1 * time.Minute)

type Claims struct {
	Username uint `json:"username"`
	jwt.StandardClaims
}

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	LoginAuth := &models.Login{}
	utils.ParsBody(r, LoginAuth)
	b, _ := LoginAuth.Login()
	errRespose := map[string]string{}
	if b == nil {
		w.WriteHeader(http.StatusUnauthorized)
		errRespose["message"] = "Unauthorized"
		res, _ := json.Marshal(errRespose)
		w.Write(res)
		return
	}
	claims := &Claims{
		Username: b.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errRespose["message"] = "Internal Sercer Error"
		res, _ := json.Marshal(errRespose)
		w.Write(res)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
	succsesRespond := map[string]string{}
	succsesRespond["message"] = "success"
	res, _ := json.Marshal(succsesRespond)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func AuthRegister(w http.ResponseWriter, r *http.Request) {
	register := &models.User{}
	utils.ParsBody(r, register)
	b := register.Register()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

type Body struct {
	Email string `json:"email"`
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	errRespose := map[string]string{}
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			errRespose["message"] = "Unauthorized"
			w.Write([]byte(errRespose["message"]))
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tokenStr := cookie.Value
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			errRespose["message"] = "Unauthorized 2"
			w.Write([]byte(errRespose["message"]))

			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		errRespose["message"] = "Unauthorized token"
		w.Write([]byte(errRespose["message"]))

		return
	}
	fmt.Println(time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))

	if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		// fmt.Println(time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(errRespose["message"]))
}
func PassworRecovery(w http.ResponseWriter, r *http.Request) {
	var body = &Body{}
	request := json.NewDecoder(r.Body)
	err := request.Decode(&body)
	if err != nil || body.Email == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// check if  email exist in DB
	h, user := models.EmailValidate(body.Email)
	if h == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	errRespose := map[string]string{}
	claims := &Claims{
		Username: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errRespose["message"] = "Internal Sercer Error"
		res, _ := json.Marshal(errRespose)
		w.Write(res)
		return
	}

	fmt.Println(tokenString)
	// http.SetCookie(w, &http.Cookie{
	// 	Name:    "token",
	// 	Value:   tokenString,
	// 	Expires: expirationTime,
	// })
	succsesRespond := map[string]string{}
	succsesRespond["message"] = "success"
	res, _ := json.Marshal(succsesRespond)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// auth middleware

// Example of router middlewares
// Usage: wrap handler with auth func
func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenStr := cookie.Value
		claims := &Claims{}
		fmt.Println(claims)
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			w.WriteHeader(http.StatusUnauthorized)
			return jwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		HandlerFunc(w, r)
	}
}
