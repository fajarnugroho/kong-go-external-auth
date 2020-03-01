FROM golang:alpine as builder

RUN apk add --no-cache git gcc libc-dev
RUN go get github.com/Kong/go-pluginserver

RUN mkdir /go-plugins
COPY go-extauth.go /go-plugins/go-extauth.go
RUN go build -buildmode plugin -o /go-plugins/go-extauth.so /go-plugins/go-extauth.go

FROM kong:2.0.1-alpine

USER root
COPY --from=builder /go/bin/go-pluginserver /usr/local/bin/go-pluginserver
RUN mkdir /go-plugins
COPY --from=builder /go-plugins/go-extauth.so /go-plugins/go-extauth.so

RUN chmod -R 775 /go-plugins
USER kong
