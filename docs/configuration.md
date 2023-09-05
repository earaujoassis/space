# Application (server) configuration

Application configuration is done through the following options, which are loaded or
attempted respecting the presented order:

1. it attempts to load configuration from the `.config.local.json` file. A template
for that file is available at [configs/.config.local.sample.json](/configs/.config.local.sample.json).
The `.config.local.json` should be available at the current working directory when running
the application binary;

2. it checks for the `.config.yml` file, at the current working directory when running
the application binary, which provides information on how to connect to a
[Vault](https://www.vaultproject.io/) and loads the configuration from it. The JSON
configuration entry in Vault should follow the template at
[configs/.config.local.sample.json](/configs/.config.local.sample.json); finally,

3. and 4., it loads the `.env` file at the current working directory when running the
application binary and attempts to retrieve the configuration files from the environment
variables (either available through `.env` or through the process environment).

If the application fails to load the configuration from any of the mentioned sources, it
panics and exit the process.

## Configuration options for the application

The following options (represented by environment variables) are available for the application.
You may set them through the described configuration channels (above).

- `SPACE_ENV`

  This configuration is only obtained through an environment variable and ideally it should
  have one of the following values: `development`, `testing`, `production`. This is set by
  default to `production` in the Dockerfile.

- `SPACE_APPLICATION_KEY`

  Application key used for confirming some administrative functions at the UI. For instance,
  when escalating a user from member to admin.

- `SPACE_DATASTORE_HOST`

  The host to the RDBMS (Postgres).

- `SPACE_DATASTORE_PORT`

  The port to the RDBMS (Postgres).

- `SPACE_DATASTORE_NAME_PREFIX`

  The name prefix for the relational database. The final name has the `SPACE_ENV` as suffix.
  So, if `SPACE_ENV=production` and `SPACE_DATASTORE_NAME_PREFIX=space`, the relational
  database should be named `space_production`.

- `SPACE_DATASTORE_USER`

  The user/owner of the relational database.

- `SPACE_DATASTORE_PASSWORD`

  The password for the user/owner of the relational database.

- `SPACE_DATASTORE_SSL_MODE`

  Set the SSL Mode when connecting to the Postgres RDBMS. You may refer to Postgres documentation
  on that: [postgresql.org/docs/current/libpq-ssl.html#LIBPQ-SSL-PROTECTION](https://www.postgresql.org/docs/current/libpq-ssl.html#LIBPQ-SSL-PROTECTION).

- `SPACE_MAIL_FROM`

  Email address used in the `From` field when sending e-mail messages.

- `SPACE_MAILER_ACCESS`

  The application uses AWS SES by default to send e-mail messages. You should provide credentials
  to access AWS SES using the following format: `"AccessKeyId:SecretAccessKey:Region"`.

- `SPACE_MEMORY_STORE_HOST`

  The host address to the Redis instance.

- `SPACE_MEMORY_STORE_PORT`

  The port to the Redis instance.

- `SPACE_MEMORY_STORE_INDEX`

  The index of the memory store in Redis.

- `SPACE_MEMORY_STORE_PASSWORD`

  The password used in Redis, through the `AUTH` command.

- `SPACE_SESSION_SECRET`

  The Cookie Session secret for the web application.

- `SPACE_SESSION_SECURE`

  A boolean (`true` | `false`) to set the Secure Attribute of a Session Cookie. Please refer to following resource:
  [owasp.org/www-community/controls/SecureCookieAttribute](https://owasp.org/www-community/controls/SecureCookieAttribute)

- `SPACE_STORAGE_SECRET`

  Key used to encrypt sensitive data in the relational database.

## Environment variables set in the Dockerfile

The following environment variables are set in the Dockerfile and could impact in the
application (mostly in the build process):

- `NODE_ENV=production`: for building and bundling web assets
- `GIN_MODE=release`: the mode which Gin-Gonic is executed

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
