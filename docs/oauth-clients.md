# OAuth Clients

## Ruby's `oauth2` gem

The following example could be used to setup the [`oauth2`](https://rubygems.org/gems/oauth2) rubygem to connect to the `earaujoassis/space` OAuth2 provider:

```rb
require 'oauth2'
require 'base64'

client_key = "<application's CLIENT_KEY, provided by space web UI>"
client_secret = "<application's CLIENT_SECRET, provided by space web UI>"
base_url = '<space server base url, like: http://localhost:9000>'

# `scope` must be `read` if you'd like to read user's content, like `first_name` and `last_name`. Also, the application must have the `read` scope --- you may set that through the web UI.
client = OAuth2::Client.new(client_key, client_secret, site: base_url, scope: 'read')

application_callback_url = "<as configured in space's web UI>"

authorize_url = client.auth_code.authorize_url(redirect_uri: application_callback_url, scope: 'read')

# Send a redirect to `authorize_url` to request for the authorization code grant.

code = "<obtained from a redirect by space, once the user authorizes the application>"
authorization = Base64.strict_encode64("#{client_key}:#{client_secret}").strip

token = client.auth_code.get_token(
  code,
  redirect_uri: application_callback_url,
  headers: { 'Authorization' => "Basic #{authorization}" }
)

response = token.post('/api/users/introspect', body: { user_id: token.params["user_id"] })

# `response.body` will contain a JSON string with user's data, like `first_name` and `last_name`.
```

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
