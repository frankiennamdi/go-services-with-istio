# Sample Rest Microservice In Go

## Requirement 
* Go 1.11
* Make
* Istio 0.6.0

## Services 

### Patient Details Service

A service that returns the details of a patient. 

### Patient Service

A service that returns information about the patient which includes data from the details service. 

## Build and Deploy 
make patientservice test build deploy-local
make patientdetailsservice test build deploy-local

## Improvements

* upgrade istio 
* implement persistence 
* more testing

