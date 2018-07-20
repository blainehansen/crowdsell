FROM golang:1.10.0-alpine3.7

RUN apk add --no-cache bash openssh build-base alpine-sdk

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

RUN mv `which dep` /usr/bin

WORKDIR $GOPATH/src/main

RUN go get github.com/cespare/reflex
RUN go install github.com/cespare/reflex

RUN go get gopkg.in/src-d/go-kallax.v1/...
RUN go install gopkg.in/src-d/go-kallax.v1/...

COPY ./server .
