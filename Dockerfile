# Dockerfile for store-service
FROM prom/busybox:glibc
MAINTAINER Caihua Yin <alend.yin@gmail.com>

COPY bin/store-service /opt/store-service/bin/
COPY config.base.yaml /opt/store-service/

WORKDIR /opt/store-service

ENTRYPOINT ["/opt/store-service/bin/store-service"]

EXPOSE 8001
