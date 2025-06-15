package integration

import (
	"github.com/earaujoassis/space/test/factory"
)

func (s *OAuthProviderSuite) TestTokenGrant() {
	var code string
	client := factory.NewClient()
	user := factory.NewUser()

	s.Run("should have a valid cookie and code", func() {
		s.Client.StartSession(user)

		s.True(s.Client.HasSessionCookie(), "Session cookie should be set")

		response := s.Client.PostAuthorize(Code, client.Key, "http://localhost/callback", "test-state", true)

		s.Equal(302, response.StatusCode)
		s.True(response.HasKeyInQuery("code"))

		code = response.Query["code"]
	})

	s.Run("should return error if redirect URI has changed", func() {
		response := s.Client.PostTokenComplete(client.BasicAuthEncode(), AuthorizationCode, code, "http://localhost/another/callback")
		json := response.JSON

		s.Equal(400, response.StatusCode)
		s.Equal("invalid_grant", json["error"])
	})

	s.Run("should return error if attempting to retrieve session token through another client", func() {
		second_client := factory.NewClient()
		response := s.Client.PostTokenComplete(second_client.BasicAuthEncode(), AuthorizationCode, code, "http://localhost/callback")
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
	})
}
