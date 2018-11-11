# Sample Rest Microservice In Go

## Requirement 
* Go 1.11
* Make
* Istio 0.6.0
* Docker for Mac Edge / With Kubernettes 

## Services 

### Patient Details Service

A service that returns the details of a patient. 

#### test, build and deploy 
`make test build deploy-local`

### Patient Service

A service that returns information about the patient which includes data from the details service. 

## Build and Deploy 
make patientservice test build deploy-local
make patientdetailsservice test build deploy-local

## Test 

`curl http://localhost/data-service/patients/2`

## Improvements

* upgrade istio 
* implement persistence 
* more testing

