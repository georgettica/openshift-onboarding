FROM alpine:latest

LABEL maintainer="Rick Rackow <rrackow@redhat.com>"

RUN apk add --no-cache \
    curl \
    git \
    openssh-client \
    rsync \
    build-base \
    libc6-compat

RUN mkdir -p /usr/local/src && \
    cd /usr/local/src && \
    curl -L https://github.com/gohugoio/hugo/releases/download/v0.59.1/hugo_extended_0.59.1_Linux-64bit.tar.gz | tar -xz && \
    mv hugo /usr/local/bin/hugo && \
    addgroup -Sg 1000 hugo && \
    adduser -Sg hugo -u 1000 -h /src hugo

WORKDIR /src

EXPOSE 1313
