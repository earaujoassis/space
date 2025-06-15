package integration

func (s *OAuthProviderSuite) TestTokenRefreshGrant() {
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

	s.Run("should return error if attempting to refresh session token through another client", func() {
		second_client := s.Factory.NewClient()
		response := s.Client.PostTokenRefresh(second_client.BasicAuthEncode(), refreshToken, "public")
		json := response.JSON

		s.Equal(400, response.StatusCode)
		s.Equal("invalid_grant", json["error"])
	})

	s.Run("should return error if attempting to refresh session token with the access session token", func() {
		response := s.Client.PostTokenRefresh(client.BasicAuthEncode(), accessToken, "public")
		json := response.JSON

		s.Equal(400, response.StatusCode)
		s.Equal("invalid_grant", json["error"])
	})

	s.Run("should return success if all parameters are correct", func() {
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

		response = s.Client.PostTokenRefresh(client.BasicAuthEncode(), refreshToken, "public")
		json = response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal("Bearer", json["token_type"])
		s.True(response.HasKeyInJSON("access_token"))
		s.True(response.HasKeyInJSON("refresh_token"))
		s.True(response.HasKeyInJSON("expires_in"))
	})
}
