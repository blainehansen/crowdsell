FROM golang:1.10.0-alpine3.7

RUN apk add --no-cache curl
RUN apk add --no-cache bash

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR $GOPATH/src/crowd-sell
