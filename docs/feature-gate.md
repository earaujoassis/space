# Feature Gate

Space uses a feature-gate to turn on or off some critical functionalities across the whole application. It relies on a memory store for that and it works on a toggle-based behaviour: when the key-field exists, the feature is available; when the key-field is not available at the memory store, the feature is not available.

Thus by default all critical features are unavailable.

## Feature toggle

Use the following command to toggle a feature-flag:

```sh
$ go run main.go feature
```

or

```sh
$ go build -o cmd/space main.go
$ cmd/space feature
```

It will request for the feature key (as described below).

## Available features

- `user.create`: it turns on the sign-up option throughout the entire application;
- `user.adminify`: it turns on the option to make a given user (post-sign-in) to turn her/himself an admin, given the provided application-key for that.

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
