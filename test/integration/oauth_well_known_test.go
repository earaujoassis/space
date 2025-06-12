package integration

import (
	"fmt"
	"strings"
)

func interfaceSliceToStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		if str, ok := v.(string); ok {
			result[i] = str
		} else {
			result[i] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func (s *OAuthProviderSuite) TestWellKnownMetadata() {
	s.Run("should return a valid well-known metadata", func() {
		response := s.Client.GetMetadata()
		json := response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal("openid public read write", strings.Join(interfaceSliceToStringSlice(json["scopes_supported"].([]interface{})), " "))
		s.Equal("code", strings.Join(interfaceSliceToStringSlice(json["response_types_supported"].([]interface{})), " "))
		s.Equal("query", strings.Join(interfaceSliceToStringSlice(json["response_modes_supported"].([]interface{})), " "))
		s.Equal("authorization_code", strings.Join(interfaceSliceToStringSlice(json["grant_types_supported"].([]interface{})), " "))
		s.Equal("client_secret_basic", strings.Join(interfaceSliceToStringSlice(json["token_endpoint_auth_methods_supported"].([]interface{})), " "))
		s.Equal("https://github.com/earaujoassis/space", json["service_documentation"])
		s.Equal("en-US", strings.Join(interfaceSliceToStringSlice(json["ui_locales_supported"].([]interface{})), " "))
		s.Equal("client_secret_basic", strings.Join(interfaceSliceToStringSlice(json["revocation_endpoint_auth_methods_supported"].([]interface{})), " "))
		s.Equal("client_secret_basic", strings.Join(interfaceSliceToStringSlice(json["introspection_endpoint_auth_methods_supported"].([]interface{})), " "))
	})
}
