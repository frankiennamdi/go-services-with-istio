package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/frankiennamdi/go-services-with-istio/patientdetailsservice/routers"
)

// Entry point for the program
func main() {
	// Get the mux router object
	router := router.InitRoutes()
	port := 3000
	log.Printf("Service Starting on Port : %d ...", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), router); err != nil {
		log.Fatal(err)
	}
}
