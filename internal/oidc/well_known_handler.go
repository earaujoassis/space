package oidc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func getOpenIdConfigurationDefinitions(c *gin.Context) utils.H {
	baseURL := shared.GetBaseUrl(c)

	return utils.H{
		"issuer":                 baseURL,
		"authorization_endpoint": fmt.Sprintf("%s%s", baseURL, "/oidc/authorize"),
		"token_endpoint":         fmt.Sprintf("%s%s", baseURL, "/oidc/token"),
		"userinfo_endpoint":      fmt.Sprintf("%s%s", baseURL, "/oidc/userinfo"),
		"jwks_uri":               fmt.Sprintf("%s%s", baseURL, "/oidc/jwks"),
		// RECOMMENDED. URL of the OP's Dynamic Client Registration Endpoint
		// [OpenID.Registration], which MUST use the https scheme.
		// "registration_endpoint": ?
		"scopes_supported":         []string{"openid", "profile", "public", "read", "write"},
		"response_types_supported": []string{"code", "id_token", "code id_token"},
		"response_modes_supported": []string{"query", "fragment", "form_post"},
		"grant_types_supported":    []string{"authorization_code"},
		// OPTIONAL. JSON array containing a list of the Authentication
		// Context Class References that this OP supports.
		// "acr_values_supported": ?
		"subject_types_supported":               []string{"public"},
		"id_token_signing_alg_values_supported": []string{"RS256"},
		// OPTIONAL. JSON array containing a list of the JWE encryption
		// algorithms (alg values) [JWT] supported by the OP for the ID
		// Token to encode the Claims in a JWT .
		// "id_token_encryption_alg_values_supported": ?
		// OPTIONAL. JSON array containing a list of the JWE encryption
		// algorithms (enc values) [JWT] supported by the OP for the ID
		// Token to encode the Claims in a JWT .
		// "id_token_encryption_enc_values_supported": ?
		// OPTIONAL. JSON array containing a list of the JWS [JWS] signing
		// algorithms (alg values) [JWA] [JWT] supported by the UserInfo
		// Endpoint to encode the Claims in a JWT . The value none MAY be
		// included.
		// "userinfo_signing_alg_values_supported": ?
		// OPTIONAL. JSON array containing a list of the JWE [JWE] encryption
		// algorithms (alg values) [JWA] [JWT] supported by the UserInfo
		// Endpoint to encode the Claims in a JWT .
		// "userinfo_encryption_alg_values_supported": ?
		// OPTIONAL. JSON array containing a list of the JWE encryption
		// algorithms (enc values) [JWA] [JWT] supported by the UserInfo
		// Endpoint to encode the Claims in a JWT .
		// "userinfo_encryption_enc_values_supported": ?
		// OPTIONAL. JSON array containing a list of the JWS signing
		// algorithms (alg values) supported by the OP for Request Objects,
		// which are described in Section 6.1 of OpenID Connect Core 1.0
		// [OpenID.Core]. These algorithms are used both when the Request
		// Object is passed by value (using the request parameter) and
		// when it is passed by reference (using the request_uri parameter).
		// Servers SHOULD support none and RS256.
		// "request_object_signing_alg_values_supported": ?
		// OPTIONAL. JSON array containing a list of the JWE encryption
		// algorithms (alg values) supported by the OP for Request Objects.
		// These algorithms are used both when the Request Object is passed
		// by value and when it is passed by reference.
		// "request_object_encryption_alg_values_supported": ?
		// OPTIONAL. JSON array containing a list of the JWE encryption
		// algorithms (enc values) supported by the OP for Request Objects.
		// These algorithms are used both when the Request Object is passed
		// by value and when it is passed by reference.
		// "request_object_encryption_enc_values_supported": ?
		"token_endpoint_auth_methods_supported":            []string{"client_secret_basic"},
		"token_endpoint_auth_signing_alg_values_supported": []string{"RS256"},
		// OPTIONAL. JSON array containing a list of the display parameter
		// values that the OpenID Provider supports. These values are
		// described in Section 3.1.2.1 of OpenID Connect Core 1.0 [OpenID.Core].
		// "display_values_supported": ?
		"claim_types_supported": []string{"normal"},
		"claims_supported": []string{
			"sub", "name", "given_name", "family_name", "preferred_username", "zoneinfo", "locale", "updated_at",
		},
		"service_documentation":            "https://github.com/earaujoassis/space",
		"claims_locales_supported":         []string{"en-US"},
		"ui_locales_supported":             []string{"en-US"},
		"claims_parameter_supported":       false,
		"request_parameter_supported":      false,
		"request_uri_parameter_supported":  true,
		"require_request_uri_registration": false,
		"op_policy_uri":                    "https://quatrolabs.com/privacy-policy",
		"op_tos_uri":                       "https://quatrolabs.com/terms-of-service",
	}
}

func getOpenIdConfiguration(c *gin.Context) {
	c.Header("Cache-Control", "public, max-age=3600")
	c.JSON(http.StatusOK, getOpenIdConfigurationDefinitions(c))
}
