FROM golang:alpine as builder

RUN mkdir -p /web-services/identity-service
WORKDIR /web-services
COPY go.mod .
COPY go.sum .
COPY proto proto

RUN apk update && apk add --update --no-cache ca-certificates git

RUN go mod download

WORKDIR /web-services/identity-service
COPY identity-service/. .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build 

FROM alpine:latest

RUN mkdir /identity-service
WORKDIR /identity-service

COPY --from=builder /web-services/identity-service/identity-service .

CMD ["./identity-service"]