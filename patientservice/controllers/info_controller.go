package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/frankiennamdi/go-services-with-istio/patientservice/config"
	"github.com/frankiennamdi/go-services-with-istio/patientservice/models"
	"github.com/gorilla/mux"
)

type Handler struct {
	Config *config.Config
	Handle func(e *config.Config, w http.ResponseWriter, r *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Handle(h.Config, w, r)
}

func Status(w http.ResponseWriter, r *http.Request) {
	log.Println("Service is up")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("up"))
}

// Handler for HTTP Get - "/data-service/patients/{id}"
func Details(config *config.Config, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientId, err := strconv.Atoi(vars["id"])
	url := fmt.Sprintf("%s/patient/%d", config.PatientDetailServiceUrl, patientId)
	log.Printf("Retrieving detail for patient %s", url)

	response, err := http.Get(url)
	if err != nil {
		log.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "cannot obtain patient details")
		return
	}
	log.Printf("Retrieving detail for patient %d with status %s", patientId, response.Status)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err.Error())
		respondWithError(w, http.StatusBadRequest, "cannot read patient details")
		return
	}

	var patientDetail models.PatientDetail
	json.Unmarshal(responseData, &patientDetail)

	var patientInfo models.PatientInfo
	patientInfo.PatientDetail = patientDetail

	resp, err := json.Marshal(&patientInfo)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Cannot create patient info")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
