FROM earaujoassis/golang-node:latest
MAINTAINER Ewerton Assis <earaujoassis@gmail.com>

ENV ENVIRONMENT production
ENV NODE_ENV production
ENV GIN_MODE release
COPY . /go/src/github.com/earaujoassis/space
WORKDIR /go/src/github.com/earaujoassis/space
ENTRYPOINT goreman start
