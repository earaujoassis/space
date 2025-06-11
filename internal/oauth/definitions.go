package oauth

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	// RFC 6749 Error types

	// InvalidRequest error type
	InvalidRequest string = "invalid_request"
	// UnauthorizedClient error type
	UnauthorizedClient string = "unauthorized_client"
	// AccessDenied error type
	AccessDenied string = "access_denied"
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

	// OpenID Connect / Extensions Error types

	// InvalidSession error type
	InvalidSession string = "invalid_session"

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

func getBaseUrl(c *gin.Context) string {
	scheme := "http"
    if c.Request.TLS != nil {
        scheme = "https"
    }
	return fmt.Sprintf("%s://%s", scheme, c.Request.Host)
}
