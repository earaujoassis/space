package integration

func (s *OAuthProviderSuite) TestTokenRevoke() {
	var accessToken, refreshToken string
	client := s.Factory.NewClient()
	user := s.Factory.NewUser()

	s.Run("should successfully retrieve tokens", func() {
		s.Client.StartSession(user)

		s.True(s.Client.HasSessionCookie(), "Session cookie should be set")

		response := s.Client.PostAuthorize(Code, client.Key, "http://localhost/callback", "test-state", true)

		s.Equal(302, response.StatusCode)
		s.True(response.HasKeyInQuery("code"))

		code := response.Query["code"]
		response = s.Client.PostTokenComplete(client.BasicAuthEncode(), AuthorizationCode, code, "http://localhost/callback")
		json := response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal("Bearer", json["token_type"])
		s.True(response.HasKeyInJSON("access_token"))
		s.True(response.HasKeyInJSON("refresh_token"))
		s.True(response.HasKeyInJSON("expires_in"))

		accessToken = json["access_token"].(string)
		refreshToken = json["refresh_token"].(string)
	})

	s.Run("should return 200 Ok if attempting to revoke token through another client", func() {
		second_client := s.Factory.NewClient()
		response := s.Client.PostRevoke(second_client.BasicAuthEncode(), accessToken)

		s.Equal(200, response.StatusCode)

		response = s.Client.PostRevoke(second_client.BasicAuthEncode(), refreshToken)

		s.Equal(200, response.StatusCode)
	})

	s.Run("should return 200 Ok if all parameters are correct", func() {
		response := s.Client.PostAuthorize(Code, client.Key, "http://localhost/callback", "test-state", true)

		s.Equal(302, response.StatusCode)
		s.True(response.HasKeyInQuery("code"))

		localCode := response.Query["code"]
		response = s.Client.PostTokenComplete(client.BasicAuthEncode(), AuthorizationCode, localCode, "http://localhost/callback")
		json := response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal("Bearer", json["token_type"])
		s.True(response.HasKeyInJSON("access_token"))
		s.True(response.HasKeyInJSON("refresh_token"))
		s.True(response.HasKeyInJSON("expires_in"))

		accessToken = json["access_token"].(string)
		refreshToken = json["refresh_token"].(string)

		response = s.Client.PostRevoke(client.BasicAuthEncode(), accessToken)

		s.Equal(200, response.StatusCode)

		response = s.Client.PostRevoke(client.BasicAuthEncode(), refreshToken)

		s.Equal(200, response.StatusCode)
	})
}
