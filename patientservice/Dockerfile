FROM golang:1.11.0-alpine as builder
RUN apk add git
COPY ./patientservice/go.mod ./patientservice/go.sum /
COPY ./patientservice /patientservice
WORKDIR /
RUN go mod download
RUN CGO_ENABLED=0  go build -o /app/patientservice

FROM alpine as runner
RUN apk add --update \
    curl \
    && rm -rf /var/cache/apk/*
COPY --from=builder /app/patientservice /app/patientservice
ENTRYPOINT /app/patientservice
EXPOSE 2000

