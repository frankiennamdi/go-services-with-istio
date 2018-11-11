package routers

import (
	"github.com/frankiennamdi/go-services-with-istio/patientservice/config"
	"github.com/frankiennamdi/go-services-with-istio/patientservice/controllers"
	"github.com/gorilla/mux"
)

func setRoutes(config *config.Config, router *mux.Router) *mux.Router {
	httpHandler := controllers.Handler{Config: config, Handle: controllers.Details}
	router.Handle("/data-service/patients/{id}", httpHandler).Methods("GET")
	router.HandleFunc("/data-service/healthcheck/status", controllers.Status).Methods("GET")
	return router
}

func InitRoutes(config *config.Config) *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = setRoutes(config, router)
	return router
}
