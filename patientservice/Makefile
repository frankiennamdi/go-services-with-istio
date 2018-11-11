
.PHONY: test build deploy
.DEFAULT_GOAL := test
test:
	go test ./... -v

build:
	dep ensure && docker build --no-cache -t frankiennamdi/patientservice .

deploy-local: SHELL := /bin/bash
deploy-local:
	kubectl apply -f <(istioctl kube-inject -f "./kube/deployment-istio.yml")

undeploy-local: 
	kubectl delete -f kube/deployment-istio.yml

view-local:
	kubectl get no,po,deployments,svc
	