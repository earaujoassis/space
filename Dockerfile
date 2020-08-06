FROM golang:1.12.12-alpine3.10

LABEL "com.quatrolabs.space"="quatroLABS Space"
LABEL "description"="A user management microservice; OAuth 2 provider"

RUN apk add --update --no-cache \
    binutils-gold \
    curl \
    g++ \
    gcc \
    gnupg \
    libgcc \
    linux-headers \
    make \
    python \
    postgresql \
    postgresql-contrib \
    postgresql-libs \
    postgresql-dev \
    git

RUN apk add --update --no-cache nodejs
RUN apk add --update --no-cache yarn

ENV PATH=/usr/local/bin:$PATH

ENV PORT=9000
ENV NODE_ENV=production
ENV GIN_MODE=release
ENV SPACE_ENV=production
ENV GO111MODULE=on

RUN mkdir -p /app

WORKDIR /app

COPY . /app

RUN yarn install && yarn build

EXPOSE 9000

ENTRYPOINT [ "go", "run", "main.go" ]
CMD [ "serve" ]
