FROM golang:1-stretch

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

WORKDIR $GOPATH/src/main

RUN go get github.com/cespare/reflex
