FROM golang:1.8.3
MAINTAINER Ewerton Assis <earaujoassis@gmail.com>

ENV NODE_ENV production
ENV GIN_MODE release
ENV ENV production
RUN go get github.com/tools/godep
RUN go get github.com/mattn/goreman
RUN mkdir -p /go/src
RUN mkdir -p /go/src/github.com
RUN mkdir -p /go/src/github.com/earaujoassis
COPY . /go/src/github.com/earaujoassis/space
WORKDIR /go/src/github.com/earaujoassis/space
RUN godep restore
EXPOSE 8080
CMD [ "goreman", "start" ]
