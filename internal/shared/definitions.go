package shared

const (
	CookieSessionKey = "application_token"
)

const (
	// Definitions shared between:
	// - OAuth
	// - OpenID Connect
	// - Web API

	// AccessDenied error type
	AccessDenied string = "access_denied"

	// ErrorQueryURI defines the query string for the callback redirect
	ErrorQueryURI string = "%s?error=%s&state=%s"
	// ErrorFragmentURI defines the fragment string for the callback redirect
	ErrorFragmentURI string = "%s#error=%s&state=%s"

	// QueryResponseType defines the response_mode for query
	QueryResponseType string = "query"
	// FragmentResponseType defines the response_mode for fragment
	FragmentResponseType string = "fragment"
	// FormPostReponseType defines the response_mode for form_post
	FormPostReponseType string = "form_post"
)

const (
	// OpenID Connect Error types

	// InvalidSession error type
	InvalidSession string = "invalid_session"
	// InvalidSession error type
	InvalidToken string = "invalid_token"
	// InvalidClient error type
	InvalidClient string = "invalid_client"
	// InsufficientScope error type
	InsufficientScope string = "insufficient_scope"
	// InvalidResponseMode error type
	InvalidResponseMode string = "invalid_response_mode"

	// OpenID Connect Response types

	// IDToken response type
	IDToken string = "id_token"
	// CodeIDToken response type
	CodeIDToken string = "code id_token"
	// IDTokenToken response type
	IDTokenToken string = "id_token token"
	// CodeToken response type
	CodeToken string = "code token"
	// CodeToken response type
	CodeIDTokenToken string = "code id_token token"
)

const (
	// RFC 6749 Error types

	// InvalidRequest error type
	InvalidRequest string = "invalid_request"
	// UnauthorizedClient error type
	UnauthorizedClient string = "unauthorized_client"
	// UnsupportedResponseType error type
	UnsupportedResponseType string = "unsupported_response_type"
	// InvalidScope error type
	InvalidScope string = "invalid_scope"
	// ServerError error type
	ServerError string = "server_error"
	// TemporarilyUnavailable error type
	TemporarilyUnavailable string = "temporarily_unavailable"
	// UnsupportedGrantType error type
	UnsupportedGrantType string = "unsupported_grant_type"
	// InvalidGrant error type
	InvalidGrant string = "invalid_grant"

	// RFC 6749 Grant types

	// AuthorizationCode grant type
	AuthorizationCode string = "authorization_code"
	// RefreshToken grant type
	RefreshToken string = "refresh_token"
	// Password grant type
	Password string = "password"
	// ClientCredentials grant type
	ClientCredentials string = "client_credentials"

	// RFC 6749 Response types

	// Code response type
	Code string = "code"
	// Token response type
	Token string = "token"
)

type TokenType int8

const (
	// TokenTypeIDToken
	TokenTypeIDToken TokenType = iota
	// TokenTypeAccessToken
	TokenTypeAccessToken
)
