#!/bin/sh

VERSION="us-central1-docker.pkg.dev/atsgo-340504/ats/atsserver:0.0.9";

git pull && \
docker build -t $VERSION . && \
docker push $VERSION
# docker run -p 8080:8080 $VERSION;
# docker run -d -p 8080:8080 $VERSION;
