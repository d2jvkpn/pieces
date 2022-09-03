#! /usr/bin/env bash
set -eu -o pipefail
_wd=$(pwd)
_path=$(dirname $0 | xargs -i readlink -f {})

for img in $(grep "^>" $0 | sed 's/>//'); do
    docker pull $img
done

docker images -q -f dangling=true | xargs -i docker rmi {}

exit
> ubuntu:22.04
> mysql:8
> grafana/grafana:main
> prom/prometheus:main
> node:16-alpine
> jaegertracing/all-in-one:latest
> redis:7-alpine
> golang:1.19-alpine
> alpine:latest
> bitnami/kafka:3.2
> bitnami/zookeeper:3.8
> otelcontribcol:latest
> rabbitmq:3-management
> mongo:5
> postgres:14
> busybox:latest
> r-base:latest
> openzipkin/zipkin:latest
> otel/opentelemetry-collector-contrib-dev:latest
> nginx:1.20-alpine
