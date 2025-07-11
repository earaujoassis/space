package integration

func (s *OAuthProviderSuite) TestTokenUnsupportedGrants() {
	client := s.Factory.NewClient()

	s.Run("should return error if requesting unsupported Resource Owner Password Credentials Grant", func() {
		response := s.Client.PostToken(client.BasicAuthEncode(), "password")
		json := response.JSON

		s.Equal(400, response.StatusCode)
		s.Equal("unsupported_grant_type", json["error"])
	})

	s.Run("should return error if requesting unsupported Client Credentials Grant", func() {
		response := s.Client.PostToken(client.BasicAuthEncode(), "client_credentials")
		json := response.JSON

		s.Equal(400, response.StatusCode)
		s.Equal("unsupported_grant_type", json["error"])
	})
}
