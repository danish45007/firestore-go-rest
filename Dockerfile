FROM golang:latest

LABEL maintainer="danish <danish45007@gmail.com>"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV URL 127.0.0.1:8000

RUN go build

CMD [ "./go-rest" ]