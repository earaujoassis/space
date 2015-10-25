# Space

> An user management microservice

## Setup

Install dependencies (soon it'll be migrated to `godep`):

```sh
$ go get -u github.com/gorilla/mux
$ go get -u github.com/lib/pq
$ go get -u gopkg.in/godo.v1/cmd/godo
```

## Running

It assumes [`godo`](https://github.com/go-godo/godo) is properly installed.

```sh
$ godo serve
```

## Issues

Please take a look at [/issues](https://github.com/earaujoassis/space/issues)

## Project Architecture & Design

This project doesn't use an ORM framework/library. It is based solely upon `database/sql`
and [`github.com/lib/pq`](https://github.com/lib/pq). For more details why I've chosen
this approach, please take a look at [hydrogen18.com/blog/golang-orms-and-why-im-still-not-using-one.html](http://www.hydrogen18.com/blog/golang-orms-and-why-im-still-not-using-one.html).

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
