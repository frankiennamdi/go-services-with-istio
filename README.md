# Sample Rest Microservice In Go

## Requirement 
* Go 1.11
* Make
* Istio 1.1.3
* Docker for Mac Edge / With Kubernettes 

## Services 

### Patient Details Service

A service that returns the details of a patient. 

#### test, build and deploy 
`make test build deploy-local`

### Patient Service

A service that returns information about the patient which includes data from the details service. 

## Build and Deploy 
kubectl apply -f http-gateway.yaml
make patientservice test build deploy-local
make patientdetailsservice test build deploy-local

## delete 
make patientservice undeploy-local
make patientdetailsservice undeploy-local
kubectl apply -f http-gateway.yaml

## Test 

`curl http://localhost/patient-service/patients/2`

`curl http://localhost/details-service/patient/1`


## Improvements
* implement persistence 
* more testing

