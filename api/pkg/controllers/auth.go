package controllers

import (
	"encoding/json"
	"fmt"
	models "github.com/Djancyp/go-rest/pkg/models"
	"github.com/Djancyp/go-rest/pkg/utils"
	"net/http"
)

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
	token, expirationTime, err := utils.CreateJwtWithClaim(b.Email)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errRespose["message"] = "Internal Sercer Error"
		res, _ := json.Marshal(errRespose)
		w.Write(res)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	succsesRespond := map[string]string{}
	succsesRespond["message"] = "success"
	res, _ := json.Marshal(succsesRespond)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

type RegisterRes struct {
	Email string `json:"email"`
}

func AuthRegister(w http.ResponseWriter, r *http.Request) {
	register := &models.User{}
	utils.ParsBody(r, register)
	b, err := register.Register()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b.Password = ""
	b.Token = ""
	returnUser := &RegisterRes{
		Email: b.Email,
	}
	res, _ := json.Marshal(returnUser)
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
	token, expirationTime, err := utils.RefreshJwt(tokenStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(errRespose["message"]))
}
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	//TODO: get user by id
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
	fmt.Println(user)
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
		resMessage := map[string]string{}
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				resMessage["message"] = "Unauthozie Request"
				w.Write([]byte(resMessage["message"]))
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			resMessage["message"] = "There is issue with this cookie"
			w.Write([]byte(resMessage["message"]))

			return
		}
		tkn, err := utils.ValidateJwt(cookie.Value)
		if err != nil || !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			resMessage["message"] = "Unauthorized"
			w.Write([]byte(resMessage["message"]))
			return
		}
		w.WriteHeader(http.StatusOK)
		HandlerFunc(w, r)
	}
}
