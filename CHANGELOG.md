# Changelog

## v0.3.0 - TBD

- Basic functionalities for the OIDC Provider
- Improvements on Client Scope selection
- Automated tests for both the OAuth and OIDC Providers

## v0.2.0 - N/A

- Basic functionality for the OAuth Provider
- Sign-up and sign-in for users are created and fully functional
- User sign-up is only available through the `user.create` feature flag
- User can become an admin by using the application key
- User can only become an admin if the `user.adminify` feature flag is enabled
- Client applications can be created by admin users
- Feature flags are enabled or disabled through the CLI command
