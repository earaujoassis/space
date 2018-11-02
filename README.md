# Space

> A user management microservice; OAuth 2 provider

Space (formely known as Jupiter) is an user authentication and authorisation
microservice intended to be used across multiple client applications. Currently,
I'm using it to provide OAuth 2 authorisation for the [earaujoassis/wallet](https://github.com/earaujoassis/wallet)
and [earaujoassis/postal](https://github.com/earaujoassis/postal) projects.
Though it is not intended to be used widely, it can be used as a reference implementation
of a Go-based OAuth 2 provider.

It uses [Gin](https://gin-gonic.github.io/gin/) as the Web Framework and [GORM](http://gorm.io/)
for Go's structure&ndash;relational design mapping. Redis is used as a memory store to keep
track of users' atomic and ephemeral actions (like the whole authorisation process under the
"Authorization Code Grant" method, described in [RFC 6749, section 4.1](https://tools.ietf.org/html/rfc6749#section-4.1)).
For the user's authentication process, it is mandatory to use Two-factor authentication (Time-based
One-time Password).

It is not planned to implement all authorisation methods described in RFC 6749 but sections
4.1 and [4.3](https://tools.ietf.org/html/rfc6749#section-4.3), the "Resource Owner Password
Credentials Grant".

## Setup & running

Space is build on top of Golang and Node.js; both are manageable on top of [`asdf`](https://github.com/asdf-vm/asdf) –
a `.tool-versions` file is already provided. Space also uses Redis and Postgres as data stores.
Please make sure to place this project inside the `$GOPATH` (a symbolic link could solve this or
simply *go getting* this like `$ go get github.com/earaujoassis/space`) and to create a `.env`
file like the `.sample.env` one. Once those requirements are met, you may run:

```sh
$ cd web && npm install && npm run build && cd ..
$ go get github.com/tools/godep
$ godep restore
$ go run main.go serve
$ open http://localhost:8080
```

## Developing & updating dependencies

If you're planning to setup it for development, ideally you should run:

```sh
$ bin/dev-setup
$ go run main.go serve
$ open http://localhost:8080
```

If any new dependency is used within the project, it should be tracked through:

```sh
$ godep save ./...
```

## Testing & Linting

For testing (plus code coverage) and linting, you could run:

```sh
$ go get -u golang.org/x/lint/golint
$ ENV=testing go test -race -coverprofile=c.out ./...
$ go tool cover -html=c.out -o coverage.html
$ golint ./...
```

## Deployment through a docker container

The following commands will create a *docker image* and create a *docker container*:

```sh
$ docker build -t earaujoassis/space .
$ docker run -d -p 8080:8080 earaujoassis/space
$ docker images --quiet --filter=dangling=true | xargs docker rmi
```

The project also provides a *docker-compose* setup, which could be setup through:

```sh
$ docker-compose up --build
```

## Issues

If you have any question, comment or an issue, please take a look at [/issues](https://github.com/earaujoassis/space/issues).

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
