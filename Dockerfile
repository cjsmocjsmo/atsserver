FROM golang:buster AS builder

RUN mkdir /go/src/atserver
WORKDIR /go/src/atserver

COPY main.go .
COPY reviews.go .
COPY estimates.go .
COPY admin.go .

COPY go.mod .
COPY go.sum .

RUN export GOPATH=/go/src/atserver
RUN go get -v /go/src/atserver
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main /go/src/atserver

FROM debian:bookworm-slim
# FROM ubuntu:22.04

RUN \
    apt-get update && \
    apt-get -y dist-upgrade && \
    apt-get -y install sqlite3 && \
    apt-get -y autoclean && \
    apt-get -y autoremove
RUN \
    mkdir /usr/share/ats_server && \
    chmod -R +rwx /usr/share/ats_server

RUN \
    touch /usr/share/ats_server/ATS.log

RUN \
    mkdir /usr/share/ats_server/static

RUN \
    touch /usr/share/ats_server/static/rev_db.tar.gz

RUN \
    touch /usr/share/ats_server/static/est_db.tar.gz

RUN \
    chmod -R +rwx /usr/share/ats_server/static

WORKDIR /usr/share/ats_server

COPY --from=builder /go/src/atserver/main .

ENV ATS_LOG_PATH=/usr/share/ats_server/ATS.log

STOPSIGNAL SIGINT
CMD ["/usr/share/ats_server/main"]