package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func ParsBody(r *http.Request, v interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, v)
	if err != nil {
		panic(err)
	}
}

type ErrMessage struct {
	Message string `json:"message"`
}

func AuthCookie(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	var resMessage ErrMessage
	if err != nil {
		if err == http.ErrNoCookie {
			resMessage.Message = "Unauthorized"
			ReturnErr(w, r, resMessage, http.StatusInternalServerError)
			return
		}
		resMessage.Message = "Bad Request"
		ReturnErr(w, r, resMessage, http.StatusBadRequest)
		return
	}
	tkn, err := ValidateJwt(cookie.Value)
	if err != nil || !tkn.Valid {
		resMessage.Message = "Unauthorized"
		ReturnErr(w, r, resMessage, http.StatusUnauthorized)
		return
	}

}
