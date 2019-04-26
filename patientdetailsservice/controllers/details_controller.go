package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/frankiennamdi/go-services-with-istio/patientdetailsservice/models"
	"github.com/gorilla/mux"
)

func Status(w http.ResponseWriter, r *http.Request) {
	log.Println("Service is up")
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("up"))
}

// Handler for HTTP Get - "/patient/{id}/details"
func Details(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientId, err := strconv.Atoi(vars["id"])
	log.Printf("Retrieving detail for patient #%d", patientId)

	resp, err := json.Marshal(models.PatientDetail{PatientName: "Franklin Ejoh",
		PatientAge:       5,
		PatientInsurance: "Blue Cross"})

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Patient ID")
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
