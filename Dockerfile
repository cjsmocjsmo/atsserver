FROM golang:bullseye AS builder

RUN mkdir /go/src/atserver
WORKDIR /go/src/atserver

COPY main.go .
COPY lib.go .

COPY go.mod .
COPY go.sum .

RUN export GOPATH=/go/src/atserver
RUN go get -v /go/src/atserver
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main /go/src/atserver

FROM debian:bullseye-slim
# FROM ubuntu:22.04

RUN \
    apt-get update && \
    apt-get -y dist-upgrade && \
    apt-get -y install sqlite3 && \
    apt-get -y autoclean && \
    apt-get -y autoremove && \
    mkdir /usr/share/ats_server && \
    chmod -R +rwx /usr/share/ats_server

WORKDIR /usr/share/ats_server

COPY --from=builder /go/src/atserver/main .

ENV ATS_LOG_PATH=/usr/share/ats_server

STOPSIGNAL SIGINT
CMD ["/usr/share/ats_server/main"]