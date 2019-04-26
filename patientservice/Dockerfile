FROM golang:1.11.0-alpine as builder
ADD . /go/src/github.com/frankiennamdi/go-services-with-istio/patientservice
RUN go install github.com/frankiennamdi/go-services-with-istio/patientservice

FROM alpine as runner
RUN apk add --no-cache curl
COPY --from=builder /go/bin/patientservice /go/bin/patientservice
CMD ./go/bin/patientservice
EXPOSE 2000