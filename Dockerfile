FROM earaujoassis/golang-node:latest
MAINTAINER Ewerton Assis <earaujoassis@gmail.com>

ENV NODE_ENV production
ENV GIN_MODE release
ENV ENV production
RUN groupadd -r space && useradd --create-home --shell /bin/bash -r -g space space
COPY . /go/src/github.com/earaujoassis/space
RUN chown -R space:space /go
USER space
WORKDIR /go/src/github.com/earaujoassis/space
ENTRYPOINT goreman start -p $PORT
