package integration

import (
	"strings"

	"github.com/earaujoassis/space/test/factory"
)

func (s *OIDCProviderSuite) TestAuthorizeIdTokenGrant() {
	user := factory.NewUser()
	client := factory.NewClientWithScopes("openid profile")

	s.Run("should have a valid cookie", func() {
		s.Client.StartSession(user)

		s.True(s.Client.HasSessionCookie(), "Session cookie should be set")
	})

	s.Run("should return error if missing response_type", func() {
		response := s.Client.GetAuthorize("", "test-client", "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "invalid_request")
	})

	s.Run("should return error if missing client_id", func() {
		response := s.Client.GetAuthorize(IdToken, "", "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "invalid_request")
	})

	s.Run("should return error if missing redirect_uri", func() {
		response := s.Client.GetAuthorize(IdToken, "test-client", "", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "invalid_request")
	})

	s.Run("should return error if redirect_uri is different from client setup", func() {
		response := s.Client.GetAuthorize(IdToken, client.Key, "http://localhost/another/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "invalid_request")
	})

	s.Run("should display consent screen with valid parameters", func() {
		response := s.Client.GetAuthorize(IdToken, client.Key, "http://localhost/callback", "test-state")

		s.Equal(200, response.StatusCode)
		s.Contains(response.Body, client.Name)
		s.Contains(response.Body, "<title>Space - Authorize</title>")
	})

	s.Run("should redirect to callback if access granted", func() {
		response := s.Client.PostAuthorize(IdToken, client.Key, "http://localhost/callback", "test-state", true)

		s.Equal(302, response.StatusCode)
		s.True(strings.HasPrefix(response.Location, "http://localhost/callback"))
		s.Equal("test-state", response.Fragment["state"])
		s.False(response.HasKeyInFragment("error"))
		s.True(response.HasKeyInFragment("id_token"))
	})

	s.Run("should redirect to index if access not granted", func() {
		response := s.Client.PostAuthorize(IdToken, client.Key, "http://localhost/callback", "test-state", false)

		s.Equal(302, response.StatusCode)
		s.True(strings.HasPrefix(response.Location, "http://localhost/callback#error="))
		s.Equal("test-state", response.Fragment["state"])
		s.Equal("access_denied", response.Fragment["error"])
		s.False(response.HasKeyInFragment("id_token"))
		s.True(response.HasKeyInFragment("error"))
	})
}
