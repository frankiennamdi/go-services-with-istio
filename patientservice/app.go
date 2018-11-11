package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/frankiennamdi/go-services-with-istio/patientservice/config"

	"github.com/frankiennamdi/go-services-with-istio/patientservice/routers"
)

func main() {
	// Get the mux router object
	config := config.Config{}
	config.Bind()
	router := routers.InitRoutes(&config)
	port := 2000
	log.Printf("Service Starting on Port : #%d ...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatal(err)
	}
}
