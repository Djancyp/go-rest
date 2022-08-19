package controllers

import (
	"encoding/json"
	"net/http"

	models "github.com/Djancyp/go-rest/pkg/models"
	"github.com/Djancyp/go-rest/pkg/utils"
	"github.com/mitchellh/mapstructure"
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
	token, expirationTime, err := utils.CreateJwtWithClaim(b.Email, b.Roles)
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
	utils.ReturnSuccses(w, r, b)
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
	resMessage := ErrMessage{}
	user := &models.User{}
	cookie, _ := r.Cookie("token")
	claims, _ := utils.GetJwtClaims(cookie.Value)
	user.Email = claims.Email
	utils.ParsBody(r, user)
	b, _ := user.UpdatePassword()
	if b == nil {
		resMessage.Message = "Unauthorized"
		utils.ReturnErr(w, r, resMessage, http.StatusUnauthorized)
		return
	}

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
	h, _ := models.EmailValidate(body.Email)
	if h == false {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	succsesRespond := map[string]string{}
	succsesRespond["message"] = "success"
	res, _ := json.Marshal(succsesRespond)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

type AuthRole struct {
	ID   uint64 `json:"id"`
	Role string `json:"role"`
}

func AddRole(w http.ResponseWriter, r *http.Request) {

}

// auth middleware

// Example of router middlewares
// Usage: wrap handler with auth func
func Auth(HandlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.AuthCookie(w, r)
		HandlerFunc(w, r)
	}
}

type Role struct {
	ID          uint64 `json:"id"`
	Role        string `json:"role"`
	Description string `json:"description"`
}

// auth type check middleware
func AuthRoles(HandlerFunc http.HandlerFunc, roles []AuthRole) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		utils.AuthCookie(w, r)
		cookie, err := r.Cookie("token")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		claims, _ := utils.GetJwtClaims(cookie.Value)
		var claims_roles []Role
		mapstructure.Decode(claims.Role, &claims_roles)
		// check if y contains role
		result := compairAuth(claims_roles, roles)
		if !result {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}
		HandlerFunc(w, r)
	}
}
func compairAuth(claims_roles []Role, roles []AuthRole) bool {
	for _, claims_role := range claims_roles {
		for _, v2 := range roles {
			if v2.Role == claims_role.Role {
				return true
			}
		}

	}

	return false
}
