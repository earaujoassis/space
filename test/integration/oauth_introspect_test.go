package integration

import (
	"github.com/earaujoassis/space/test/factory"
)

func (s *OAuthProviderSuite) TestTokenIntrospect() {
	var accessToken, refreshToken string
	client := factory.NewClient()
	clientModel := client.Model
	user := factory.NewUser()
	userModel := user.Model

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

	s.Run("should return active:false if attempting to introspect token through another client", func() {
		second_client := factory.NewClient()
		response := s.Client.PostIntrospect(second_client.BasicAuthEncode(), accessToken)
		json := response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal(false, json["active"])

		response = s.Client.PostIntrospect(second_client.BasicAuthEncode(), refreshToken)
		json = response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal(false, json["active"])
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

		response = s.Client.PostIntrospect(client.BasicAuthEncode(), accessToken)
		json = response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal(true, json["active"])
		s.Equal("Bearer", json["token_type"])
		s.Equal(clientModel.Key, json["client_id"])
		s.Equal(userModel.Username, json["username"])
		s.Equal(userModel.PublicID, json["sub"])

		response = s.Client.PostIntrospect(client.BasicAuthEncode(), refreshToken)
		json = response.JSON

		s.Equal(200, response.StatusCode)
		s.Equal(true, json["active"])
		s.Equal(nil, json["token_type"])
		s.Equal(clientModel.Key, json["client_id"])
		s.Equal(userModel.Username, json["username"])
		s.Equal(userModel.PublicID, json["sub"])
	})
}
