package utils

import (
	"encoding/json"
	"net/http"
)

func ReturnSuccses(w http.ResponseWriter, r *http.Request, value interface{}) {
	succsesRespond := value
	res, _ := json.Marshal(succsesRespond)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
func ReturnErr(w http.ResponseWriter, r *http.Request, value interface{}, status int) {
	succsesRespond := value
	res, _ := json.Marshal(succsesRespond)
	w.WriteHeader(status)
	w.Write(res)
	return
}
