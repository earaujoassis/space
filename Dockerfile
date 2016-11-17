FROM earaujoassis/golang-node:latest
MAINTAINER Ewerton Assis <earaujoassis@gmail.com>

ENV ENVIRONMENT production
ENV NODE_ENV production
ENV GIN_MODE release
RUN groupadd -r johndoe && useradd -r -g johndoe johndoe
RUN chown -R johndoe:johndoe /go
USER johndoe
COPY . /go/src/github.com/earaujoassis/space
WORKDIR /go/src/github.com/earaujoassis/space
ENTRYPOINT goreman start
