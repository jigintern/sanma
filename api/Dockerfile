FROM golang:1.14.0-alpine3.11

RUN apk add --update --no-cache ca-certificates git

RUN mkdir /go-api
WORKDIR /go-api

COPY . .
RUN go mod download

RUN CGO_ENABLED=0

CMD go run server.go
