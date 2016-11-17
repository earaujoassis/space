FROM earaujoassis/golang-node:latest
MAINTAINER Ewerton Assis <earaujoassis@gmail.com>

ENV ENVIRONMENT production
ENV NODE_ENV production
ENV GIN_MODE release
RUN groupadd -r johndoe && useradd --create-home --shell /bin/bash -r -g johndoe johndoe
COPY . /go/src/github.com/earaujoassis/space
RUN chown -R johndoe:johndoe /go
USER johndoe
WORKDIR /go/src/github.com/earaujoassis/space
ENTRYPOINT goreman start
