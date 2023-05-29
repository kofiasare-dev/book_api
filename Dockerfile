FROM golang:1.20.4-alpine

RUN apk add --no-cache && apk add --virtual build-dependencies\
    build-base \
    gcc\
    wget\
    git

RUN mkdir -p /app

WORKDIR /app

COPY go.*  ./

RUN go mod download

COPY . ./

# RUN go build -o /usr/bin/dlm
