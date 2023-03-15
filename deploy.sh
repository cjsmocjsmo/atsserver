#!/bin/sh

VERSION="atsserver:0.0.1";

docker build -t $VERSION && \
docker run -p 8080:8080 $VERSION;