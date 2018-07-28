FROM golang:1-stretch

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
RUN mv `which dep` /usr/bin

WORKDIR $GOPATH/src/main

RUN go get github.com/cespare/reflex
RUN go install github.com/cespare/reflex

COPY ./server .
