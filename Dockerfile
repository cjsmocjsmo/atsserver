FROM golang:buster AS builder

RUN mkdir /go/src/atserver
WORKDIR /go/src/atserver

COPY admin.go .
COPY estimates.go .
COPY go.mod .
COPY go.sum .
COPY main.go .
COPY mktables.go .
COPY reviews.go .

RUN export GOPATH=/go/src/atserver
RUN go get -v /go/src/atserver
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o main /go/src/atserver

FROM debian:bookworm-slim

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
    mkdir /usr/share/ats_server/static && \
    touch /usr/share/ats_server/static/rev_db.tar.gz && \
    touch /usr/share/ats_server/static/est_db.tar.gz && \
    chmod -R +rwx /usr/share/ats_server/static

RUN \
    mkdir /usr/share/ats_server/users && \
    chmod -R +rwx /usr/share/ats_server/users

WORKDIR /usr/share/ats_server

COPY --from=builder /go/src/atserver/main .

COPY user1.yaml ./users

COPY user2.yaml ./users

ENV ATS_PATH=/usr/share/ats_server
ENV ATS_LOG_PATH=/usr/share/ats_server/ATS.log
ENV ATS_DOCKER_VAR=DOCKER

STOPSIGNAL SIGINT
CMD ["/usr/share/ats_server/main"]