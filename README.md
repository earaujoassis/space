# Space [![CircleCI](https://dl.circleci.com/status-badge/img/gh/earaujoassis/space/tree/master.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/earaujoassis/space/tree/master) [![codecov](https://codecov.io/gh/earaujoassis/space/branch/master/graph/badge.svg)](https://codecov.io/gh/earaujoassis/space) [![Go Report Card](https://goreportcard.com/badge/github.com/earaujoassis/space)](https://goreportcard.com/report/github.com/earaujoassis/space)

> A user management microservice; OAuth 2.0 and OIDC provider

Space (formerly known as Jupiter) is a user authentication and authorisation microservice intended to be used across multiple client applications. Currently, it's used to provide OAuth 2 authorisation for the [earaujoassis/wallet](https://github.com/earaujoassis/wallet) and [earaujoassis/watchman](https://github.com/earaujoassis/watchman) projects. Though it is not intended to be used widely, it can be used as a reference implementation of a Golang-based OAuth 2 provider.

It uses [Gin](https://gin-gonic.github.io/gin/) as the Web Framework and [GORM](http://gorm.io/) for Golang's structure&ndash;relational mapping. Redis is used as a memory store to keep track of users' atomic and ephemeral actions (like the whole authorisation process under the "Authorization Code Grant" method, described in [RFC 6749, section 4.1](https://tools.ietf.org/html/rfc6749#section-4.1)). For the user's authentication process, it is mandatory to use Two-factor authentication (Time-based One-time Password).

Additionally, the application supports the use of Hashicorp's Vault in order to obtain configuration information and secrets. It is planned to use Vault as a way to store client keys and secrets, if the OAuth client application supports it.

It is not planned to implement all authorisation methods described in RFC 6749 but section 4.1 only, "Authorization Code Grant". As OAuth 2.1 is under draft, and following [RFC 7636](https://www.rfc-editor.org/rfc/rfc7636) "Proof Key for Code Exchange by OAuth Public Clients", [RFC 9700](https://www.ietf.org/rfc/rfc9700.html) "Best Current Practice for OAuth 2.0 Security", and the drafts for OAuth 2.1 and "OAuth 2.0 for Browser-Based Applications", the methods for "Implicit Grant", "Resource Owner Password Credentials Grant", and "Client Credentials Grant" won't be implemented.

Space is based on a set of [feature flags](docs/feature-gate.md) enabled/disabled through a Redis store.

### Implemented RFCs

- [RFC 6749](https://tools.ietf.org/html/rfc6749): The OAuth 2.0 Authorization Framework
- [RFC 6750](https://datatracker.ietf.org/doc/html/rfc6750): The OAuth 2.0 Authorization Framework: Bearer Token Usage
- [RFC 9700](https://www.ietf.org/rfc/rfc9700.html): Best Current Practice for OAuth 2.0 Security
- [RFC 7009](https://datatracker.ietf.org/doc/html/rfc7009): OAuth 2.0 Token Revocation
- [RFC 7662](https://datatracker.ietf.org/doc/html/rfc7662): OAuth 2.0 Token Introspection
- [RFC 8414](https://datatracker.ietf.org/doc/html/rfc8414): OAuth 2.0 Authorization Server Metadata

### Implemented Specifications

- [OpenID Connect Core 1.0 incorporating errata set 2](https://openid.net/specs/openid-connect-core-1_0.html)
- [OpenID Connect Discovery 1.0 incorporating errata set 2](https://openid.net/specs/openid-connect-discovery-1_0.html)
- [OAuth 2.0 Multiple Response Type Encoding Practices](https://openid.net/specs/oauth-v2-multiple-response-types-1_0.html)

## Setup & Running

Space is built on top of Golang and Node.js; both are manageable on top of [`asdf`](https://github.com/asdf-vm/asdf) – a `.tool-versions` file is already provided. Space also uses Redis and Postgres as data stores – and optionally Vault. Please make sure to place this project inside the `$GOPATH` (a symbolic link could solve this or simply *go getting* this like `$ go get github.com/earaujoassis/space`) or to set the `GO111MODULE=on` env var. A `.env` file is used to load environment information for the application – `.env` is already available. The application needs a `PORT` env var in order to setup its socket (it is automatically provided by `goreman`).

In order to setup the application, additional configuration is necessary. That could be delivered through Vault or through a local `.config.local.json` file. If you plan to use Vault, you need to fill the `.config.yml` file, according to the `.config.sample.yml`, and store a k/v secret according to the `.config.local.sample.json`. If you don't plan to use Vault, just create a `.config.local.json` file according to `.config.local.sample.json`. The application will try to load the `.config.local.json` by default, if it is available; otherwise, it will attempt to load it from Vault. At least one of these configuration schemes is necessary.

Once the configuration is complete, the following commands will run the server application locally:

```sh
$ scripts/setup.sh
$ go get github.com/mattn/goreman
$ goreman start
$ open http://localhost:9000
```

## Testing & Linting

For testing (plus code coverage) and linting, you could run:

```sh
$ scripts/setup.sh
$ task test
$ task lint && yarn lint
```

## Issues

If you have any question, comment or an issue, please take a look at [/issues](https://github.com/earaujoassis/space/issues).

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Carlos Assis

## Credits

Illustrations provided by [unDraw](https://undraw.co) - "Open-source illustrations for any idea you can imagine and create". For more details, take a look at [ATTRIBUTIONS.md](ATTRIBUTIONS.md)
