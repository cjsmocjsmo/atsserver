#!/bin/sh

VERSION="atsserver:0.0.2";

git pull && \
docker build -t $VERSION . && \
docker run -p 8080:8080 $VERSION;
# docker run -d -p 8080:8080 $VERSION;
