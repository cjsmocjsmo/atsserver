FROM golang:buster AS builder

RUN mkdir /go/src/atserver
WORKDIR /go/src/atserver

COPY admin.go .
COPY db_count.go .
COPY estimates.go .
COPY go.mod .
COPY go.sum .
COPY main.go .
COPY media.go .
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

RUN touch /usr/share/ats_server/ATS.log

RUN \
    mkdir /usr/share/ats_server/static && \
    chmod -R +rwx /usr/share/ats_server/static

RUN \
    mkdir /usr/share/ats_server/users && \
    chmod -R +rwx /usr/share/ats_server/users

WORKDIR /usr/share/ats_server

COPY --from=builder /go/src/atserver/main .

WORKDIR /usr/share/ats_server/users

COPY user1.yaml .

COPY user2.yaml .

# WORKDIR /use/share/ats_server/static

# COPY dbbackup.tar.gz .
# COPY est_db.tar.gz .
# COPY rev_db.tar.gz .

# WORKDIR /use/share/ats_server

# RUN \
#     chmod -R +rwx /use/share/ats_server/static/dbbackup.tar.gz && \
#     chmod -R +rwx /use/share/ats_server/static/est_db.tar.gz && \
#     chmod -R +rwx /use/share/ats_server/static/rev_db.tar.gz

ENV ATS_PATH=/usr/share/ats_server
ENV ATS_LOG_PATH=/usr/share/ats_server/ATS.log
ENV ATS_DOCKER_VAR=DOCKER

STOPSIGNAL SIGINT
CMD ["/usr/share/ats_server/main"]