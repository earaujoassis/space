# OAuth Clients

## Ruby's `oauth2` gem

The following example could be used to setup the [`oauth2`](https://rubygems.org/gems/oauth2) Ruby Gem to connect to the `earaujoassis/space` OAuth2 provider:

```rb
require 'oauth2'
require 'base64'

client_key = "<application's CLIENT_KEY, provided by space web UI>"
client_secret = "<application's CLIENT_SECRET, provided by space web UI>"
base_url = '<space server base url, like: http://localhost:9000>'

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
```

## License

[MIT License](http://earaujoassis.mit-license.org/) &copy; Ewerton Assis
