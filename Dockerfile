FROM golang:1.10.0-alpine3.7

RUN apk add --no-cache bash curl git openssh

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mv `which dep` /usr/bin

WORKDIR $GOPATH/src/main
