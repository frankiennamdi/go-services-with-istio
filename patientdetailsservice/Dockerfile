FROM golang:1.11.0-alpine as builder
ADD . /go/src/github.com/frankiennamdi/go-services-with-istio/patientdetailsservice
RUN go install github.com/frankiennamdi/go-services-with-istio/patientdetailsservice

FROM alpine as runner
COPY --from=builder /go/bin/patientdetailsservice /go/bin/patientdetailsservice
ENTRYPOINT ./go/bin/patientdetailsservice
EXPOSE 3000