package integration

import (
	"strings"
)

func (s *OIDCProviderSuite) TestWellKnownMetadata() {
	s.Run("should return a valid well-known metadata", func() {
		response := s.Client.GetMetadata()
		json := response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal("openid profile public read write", strings.Join(interfaceSliceToStringSlice(json["scopes_supported"].([]interface{})), " "))
		s.Equal("code id_token code id_token", strings.Join(interfaceSliceToStringSlice(json["response_types_supported"].([]interface{})), " "))
		s.Equal("query fragment form_post", strings.Join(interfaceSliceToStringSlice(json["response_modes_supported"].([]interface{})), " "))
		s.Equal("authorization_code", strings.Join(interfaceSliceToStringSlice(json["grant_types_supported"].([]interface{})), " "))
		s.Equal("public", strings.Join(interfaceSliceToStringSlice(json["subject_types_supported"].([]interface{})), " "))
		s.Equal("RS256", strings.Join(interfaceSliceToStringSlice(json["id_token_signing_alg_values_supported"].([]interface{})), " "))
		s.Equal("client_secret_basic", strings.Join(interfaceSliceToStringSlice(json["token_endpoint_auth_methods_supported"].([]interface{})), " "))
		s.Equal("RS256", strings.Join(interfaceSliceToStringSlice(json["token_endpoint_auth_signing_alg_values_supported"].([]interface{})), " "))
		s.Equal("normal", strings.Join(interfaceSliceToStringSlice(json["claim_types_supported"].([]interface{})), " "))
		s.Equal("sub name given_name family_name preferred_username zoneinfo locale updated_at", strings.Join(interfaceSliceToStringSlice(json["claims_supported"].([]interface{})), " "))
		s.Equal("https://github.com/earaujoassis/space", json["service_documentation"])
		s.Equal("en-US", strings.Join(interfaceSliceToStringSlice(json["claims_locales_supported"].([]interface{})), " "))
		s.Equal("en-US", strings.Join(interfaceSliceToStringSlice(json["ui_locales_supported"].([]interface{})), " "))
		s.Equal("https://quatrolabs.com/privacy-policy", json["op_policy_uri"])
		s.Equal("https://quatrolabs.com/terms-of-service", json["op_tos_uri"])
	})
}
