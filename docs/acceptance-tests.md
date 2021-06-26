# Acceptance tests

## Setup

```sh
$ brew install chromedriver
$ go get github.com/tools/godep
```

## Testing

```sh
$ SPACE_ENV=testing go test -cover ./...
```

## Generate new test case

```sh
$ ginkgo generate --agouti {description-file}
```

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
