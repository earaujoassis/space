package oauth

const (
    // Error types
    InvalidRequest                string = "invalid_request"
    UnauthorizedClient            string = "unauthorized_client"
    AccessDenied                  string = "access_denied"
    UnsupportedResponseType       string = "unsupported_response_type"
    InvalidScope                  string = "invalid_scope"
    ServerError                   string = "server_error"
    TemporarilyUnavailable        string = "temporarily_unavailable"
    UnsupportedGrantType          string = "unsupported_grant_type"
    InvalidGrant                  string = "invalid_grant"

    // Grant types
    AuthorizationCode             string = "authorization_code"
    RefreshToken                  string = "refresh_token"
    Password                      string = "password"
    ClientCredentials             string = "client_credentials"

    // Response types
    Code                          string = "code"
    Token                         string = "token"
)
