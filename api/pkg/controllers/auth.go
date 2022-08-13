package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	models "github.com/Djancyp/go-rest/pkg/models"
	"github.com/Djancyp/go-rest/pkg/utils"
)

type Message struct {
	ID    uint64 `json:"id"`
	Email string `json:"email"`
}
type ErrMessage struct {
	Message string `json:"message"`
}

func LoginAuth(w http.ResponseWriter, r *http.Request) {
	LoginAuth := &models.Login{}
	var errMessage ErrMessage
	utils.ParsBody(r, LoginAuth)
	b, _ := LoginAuth.Login()
	if b == nil {
		errMessage.Message = "Unauthorized"
		utils.ReturnErr(w, r, errMessage, 401)
		return
	}
	token, expirationTime, err := utils.CreateJwtWithClaim(b.Email)
	if err != nil {
		errMessage.Message = "StatusInternalServerError"
		utils.ReturnErr(w, r, errMessage, 500)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	var message Message
	message.ID = b.ID
	message.Email = b.Email
	utils.ReturnSuccses(w, r, message)
}

type RegisterRes struct {
	Email string `json:"email"`
}

func AuthRegister(w http.ResponseWriter, r *http.Request) {
	register := &models.User{}
	utils.ParsBody(r, register)
	b, err := register.Register()
	if err != nil {
		utils.ReturnErr(w, r, err, http.StatusBadRequest)
		return
	}
	returnUser := &RegisterRes{
		Email: b.Email,
	}
	utils.ReturnSuccses(w, r, returnUser)
}

type Body struct {
	Email string `json:"email"`
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	var errRespose ErrMessage
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			errRespose.Message = "Unauthorized"
			utils.ReturnErr(w, r, errRespose, http.StatusUnauthorized)
			return
		}
		errRespose.Message = "Bad Request"
		utils.ReturnErr(w, r, errRespose, http.StatusUnauthorized)
		return
	}
	tokenStr := cookie.Value
	token, expirationTime, err := utils.RefreshJwt(tokenStr)
	if err != nil {
		errRespose.Message = "Unauthorized"
		utils.ReturnErr(w, r, errRespose, http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})

	errRespose.Message = "Success"
	utils.ReturnSuccses(w, r, errRespose)
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
