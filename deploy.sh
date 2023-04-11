#!/bin/sh

VERSION="gcr.io/atsgo-340504/atsserver:0.0.8";

git pull && \
docker build -t $VERSION . && \
docker push $VERSION
# docker run -p 8080:8080 $VERSION;
# docker run -d -p 8080:8080 $VERSION;
