package routers

import (
	"github.com/frankiennamdi/go-services-with-istio/patientdetailsservice/controllers"
	"github.com/gorilla/mux"
)

func setRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/details-service/patient/{id}", controllers.Details).Methods("GET")
	return router
}

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router = setRoutes(router)
	return router
}
