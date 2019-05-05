FROM golang:1.11.0-alpine as builder
RUN apk add git
COPY ./patientdetailsservice/go.mod ./patientdetailsservice/go.sum /
COPY ./patientdetailsservice /patientdetailsservice
WORKDIR /
RUN go mod download
RUN CGO_ENABLED=0  go build -o /app/patientdetailsservice

FROM alpine as runner
RUN apk add --update \
    curl \
    && rm -rf /var/cache/apk/*
COPY --from=builder /app/patientdetailsservice /app/patientdetailsservice
ENTRYPOINT /app/patientdetailsservice
EXPOSE 3000
