# docker build -f Dockerfile-crowdsell-go -t blainehansen/crowdsell .

FROM golang:1-stretch AS builder

# RUN apk update && apk add --no-cache git

RUN groupadd -g 999 appuser && useradd -r -u 999 -g appuser appuser

COPY . $GOPATH/src/package/app/
WORKDIR $GOPATH/src/package/app/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/compiled-app


FROM scratch

COPY --from=builder /go/bin/compiled-app /app/compiled-app
COPY --from=builder /etc/passwd /etc/passwd

USER appuser

EXPOSE 5050

ENTRYPOINT ["/app/compiled-app"]
