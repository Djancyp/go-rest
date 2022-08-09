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
