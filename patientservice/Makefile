
SHELL := /bin/bash
.PHONY: test build deploy
.DEFAULT_GOAL := test

dependencies:
	go mod vendor; go mod tidy

test: dependencies
	go test ./... -v

build:
	go mod tidy && docker build --no-cache -t frankiennamdi/patientservice -f ${PWD}/patientservice/Dockerfile ${PWD}

deploy-local:
	kubectl apply -f <(istioctl kube-inject -f "./kube/deployment-istio-1.yml")

undeploy-local: 
	kubectl delete -f kube/deployment-istio-1.yml

view-local:
	kubectl get no,po,deployments,svc
	