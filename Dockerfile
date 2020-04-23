FROM golang:1.14.2-alpine3.11

WORKDIR /go/src

COPY ./src /go/src

ENV GOPATH ''

RUN apk update && apk add --no-cache git
