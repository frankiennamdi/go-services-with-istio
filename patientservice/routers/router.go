package routers

import (
	"github.com/frankiennamdi/go-services-with-istio/patientservice/config"
	"github.com/frankiennamdi/go-services-with-istio/patientservice/controllers"
	"github.com/gorilla/mux"
)

func setRoutes(router *mux.Router) *mux.Router {
	patientDetailServiceUrl := "http://patient-details-service:3000/details-service/patient"
	config := &config.Config{PatientDetailServiceUrl: patientDetailServiceUrl}
	httpHandler := controllers.Handler{Config: config, Handle: controllers.Details}
	router.Handle("/data-service/patients/{id}", httpHandler).Methods("GET")
	router.HandleFunc("/data-service/healthcheck/status", controllers.Status).Methods("GET")
	return router
}

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = setRoutes(router)
	return router
}
