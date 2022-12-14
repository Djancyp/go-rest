package controllers

import (
	"encoding/json"
	models "github.com/Djancyp/go-rest/pkg/models"
	"github.com/Djancyp/go-rest/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllExamples(w http.ResponseWriter, r *http.Request) {
	example := models.GetAllExamples()
	res, _ := json.Marshal(example)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateExample(w http.ResponseWriter, r *http.Request) {
	CreateExample := &models.Example{}
	utils.ParsBody(r, CreateExample)
	b := CreateExample.Create()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exampleId := vars["id"]
	ID, err := strconv.ParseInt(exampleId, 0, 0)
	if err != nil {
		utils.ReturnErr(w, r, err, http.StatusBadRequest)
		return
	}
	exampleDetails, _ := models.GetExampleById(ID)
	res, _ := json.Marshal(exampleDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteExample(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exampleId := vars["id"]
	ID, err := strconv.ParseInt(exampleId, 0, 0)
	if err != nil {
		utils.ReturnErr(w, r, err, http.StatusBadRequest)
		return
	}
	exampleDetails, _ := models.DeleteExampleById(ID)
	res, _ := json.Marshal(exampleDetails)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateExampleById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	exampleId := vars["id"]
	ID, err := strconv.ParseInt(exampleId, 0, 0)
	if err != nil {
		utils.ReturnErr(w, r, err, http.StatusBadRequest)
		return
	}
	updateExample := &models.Example{}
	utils.ParsBody(r, updateExample)
	b, _ := updateExample.UpdateExample(ID)
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
