package controllers

import (
	"encoding/json"
	"github.com/Djancyp/go-rest/pkg/models"
	"net/http"
)

var newExample exampleModal.Example

type Example struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetAllExamples(w http.ResponseWriter, r *http.Request) {
	example := exampleModal.GetAllExamples()
	res, _ := json.Marshal(example)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateExample(w http.ResponseWriter, r *http.Request) {
	var example = Example{ID: "2", Name: "Example2"}
	res, _ := json.Marshal(example)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
