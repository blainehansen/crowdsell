FROM golang:1.10.0-alpine3.7

RUN apk add --no-cache bash openssh build-base alpine-sdk

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mv `which dep` /usr/bin

WORKDIR $GOPATH/src/main

COPY ./server .

RUN go install .
