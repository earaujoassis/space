# Space

> An user management microservice

## Setup

Please make sure to `git clone` (or `go get`) this project inside the `$GOPATH`
in the following structure: `$GOPATH/src/github.com/earaujoassis/space`. Install
the project dependencies:

```sh
$ cd $GOPATH/src/github.com/earaujoassis/space
$ go get
```

Please install the [`godo`](https://github.com/go-godo/godo) task runner:

```sh
$ go get -u gopkg.in/godo.v1/cmd/godo
```

## Running

It assumes [`godo`](https://github.com/go-godo/godo) is properly installed.

```sh
$ godo setup server
```

## Testing

```sh
$ go test
```

## Issues

Please take a look at [/issues](https://github.com/earaujoassis/space/issues)

## Project Architecture & Design

This project doesn't use an ORM framework/library. It is based solely upon `database/sql`
and [`github.com/lib/pq`](https://github.com/lib/pq). For more details why I've chosen
this approach, please take a look at [hydrogen18.com/blog/golang-orms-and-why-im-still-not-using-one.html](http://www.hydrogen18.com/blog/golang-orms-and-why-im-still-not-using-one.html).

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
