package integration

func (s *OIDCProviderSuite) TestWellKnownMetadata() {
	s.Run("should return a valid well-known metadata", func() {
		response := s.Client.GetMetadata()
		json := response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal("openid profile public read write", jsonValueAsSingleString(json["scopes_supported"]))
		s.Equal("code id_token code id_token", jsonValueAsSingleString(json["response_types_supported"]))
		s.Equal("query fragment form_post", jsonValueAsSingleString(json["response_modes_supported"]))
		s.Equal("authorization_code", jsonValueAsSingleString(json["grant_types_supported"]))
		s.Equal("public", jsonValueAsSingleString(json["subject_types_supported"]))
		s.Equal("RS256", jsonValueAsSingleString(json["id_token_signing_alg_values_supported"]))
		s.Equal("client_secret_basic", jsonValueAsSingleString(json["token_endpoint_auth_methods_supported"]))
		s.Equal("RS256", jsonValueAsSingleString(json["token_endpoint_auth_signing_alg_values_supported"]))
		s.Equal("normal", jsonValueAsSingleString(json["claim_types_supported"]))
		s.Equal("sub name given_name family_name preferred_username zoneinfo locale updated_at", jsonValueAsSingleString(json["claims_supported"]))
		s.Equal("https://github.com/earaujoassis/space", json["service_documentation"])
		s.Equal("en-US", jsonValueAsSingleString(json["claims_locales_supported"]))
		s.Equal("en-US", jsonValueAsSingleString(json["ui_locales_supported"]))
		s.Equal("https://quatrolabs.com/privacy-policy", json["op_policy_uri"])
		s.Equal("https://quatrolabs.com/terms-of-service", json["op_tos_uri"])
	})
}
