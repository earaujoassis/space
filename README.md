# Space

> A user management microservice; OAuth 2 provider

Space (formely known as Jupiter) is an user authentication and authorisation
microservice intended to be used across multiple client applications. Currently,
I'm using it to provide OAuth 2 authorisation for the [earaujoassis/wallet](https://github.com/earaujoassis/wallet)
and [earaujoassis/postal](https://github.com/earaujoassis/postal) projects.
Though it is not intended to be used widely, it can be used as a reference implementation
of a Go-based OAuth 2 provider.

I'm using Gin as the Web Framework and GORM for Go's structures&ndash;relational
design mapping. Redis is used as a memory store to keep track of users' atomic and
ephemeral actions (like the whole authorisation process under the "Authorization Code Grant"
method, described in [RFC 6749, section 4.1](https://tools.ietf.org/html/rfc6749#section-4.1)).
For the authentication process, it is mandatory to use Two-factor authentication (Time-based One-time Password).

It is not planned to implement all authorisation methods described in RFC 6749
but sections 4.1 and [4.3](https://tools.ietf.org/html/rfc6749#section-4.3),
the "Resource Owner Password Credentials Grant".

## Issues

Please take a look at [/issues](https://github.com/earaujoassis/space/issues)

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
