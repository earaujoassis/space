package integration

func (s *OAuthProviderSuite) TestWellKnownMetadata() {
	s.Run("should return a valid well-known metadata", func() {
		response := s.Client.GetMetadata()
		json := response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal("public read", jsonValueAsSingleString(json["scopes_supported"]))
		s.Equal("code", jsonValueAsSingleString(json["response_types_supported"]))
		s.Equal("query", jsonValueAsSingleString(json["response_modes_supported"]))
		s.Equal("authorization_code", jsonValueAsSingleString(json["grant_types_supported"]))
		s.Equal("client_secret_basic", jsonValueAsSingleString(json["token_endpoint_auth_methods_supported"]))
		s.Equal("https://github.com/earaujoassis/space", json["service_documentation"])
		s.Equal("en-US", jsonValueAsSingleString(json["ui_locales_supported"]))
		s.Equal("client_secret_basic", jsonValueAsSingleString(json["revocation_endpoint_auth_methods_supported"]))
		s.Equal("client_secret_basic", jsonValueAsSingleString(json["introspection_endpoint_auth_methods_supported"]))
	})
}
