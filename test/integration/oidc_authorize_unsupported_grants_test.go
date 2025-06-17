package integration

func (s *OIDCProviderSuite) TestAuthorizeUnsupportedGrants() {
	user := s.Factory.NewUser()
	client := s.Factory.NewClientWithScopes("openid profile")

	s.Run("should have a valid cookie", func() {
		s.Client.StartSession(user)

		s.True(s.Client.HasSessionCookie(), "Session cookie should be set")
	})

	s.Run("should return error if requesting unsupported Implicit Grant", func() {
		response := s.Client.GetAuthorize(Token, client.Key, "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "unsupported_response_type")
	})

	s.Run("should return error if requesting unsupported variations", func() {
		response := s.Client.GetAuthorize("code id_token", client.Key, "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "unsupported_response_type")

		response = s.Client.GetAuthorize("id_token token", client.Key, "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "unsupported_response_type")

		response = s.Client.GetAuthorize("code token", client.Key, "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "unsupported_response_type")

		response = s.Client.GetAuthorize("code id_token token", client.Key, "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "unsupported_response_type")
	})
}
