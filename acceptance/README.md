# Space / Acceptance tests

> A user management microservice; OAuth 2 provider

## Setup

```sh
$ brew install chromedriver
$ go get github.com/tools/godep
$ godep restore
```

## Testing

```sh
$ ENV=testing go test -cover ./...
```

## Generate new test case

```sh
$ ginkgo generate --agouti {description-file}
```

## Limitations

Currently, there is no dependency management

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
