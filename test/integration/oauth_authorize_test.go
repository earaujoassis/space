package integration

import (
	"strings"

	"github.com/earaujoassis/space/test/factory"
)

func (s *OAuthProviderSuite) TestAuthorizeGrant() {
	user := factory.NewUser()
	client := factory.NewClient()

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
		response := s.Client.GetAuthorize("code", "", "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "invalid_request")
	})

	s.Run("should return error if missing redirect_uri", func() {
		response := s.Client.GetAuthorize("code", "test-client", "", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "invalid_request")
	})

	s.Run("should return error if redirect_uri is different from client setup", func() {
		response := s.Client.GetAuthorize("code", client.Key, "http://localhost/another/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "invalid_request")
	})

	s.Run("should display consent screen with valid parameters", func() {
		response := s.Client.GetAuthorize("code", client.Key, "http://localhost/callback", "test-state")

		s.Equal(200, response.StatusCode)
		s.Contains(response.Body, client.Name)
		s.Contains(response.Body, "<title>Space - Authorize</title>")
	})

	s.Run("should redirect to callback if access granted", func() {
		response := s.Client.PostAuthorize("code", client.Key, "http://localhost/callback", "test-state", true)

		s.Equal(302, response.StatusCode)
		s.True(strings.HasPrefix(response.Location, "http://localhost/callback"))
		s.Equal("test-state", response.Query["state"])
		s.False(response.HasKeyInQuery("error"))
		s.True(response.HasKeyInQuery("code"))
		s.True(response.HasKeyInQuery("scope"))
	})

	s.Run("should redirect to index if access not granted", func() {
		response := s.Client.PostAuthorize("code", client.Key, "http://localhost/callback", "test-state", false)

		s.Equal(302, response.StatusCode)
		s.True(strings.HasPrefix(response.Location, "http://localhost/callback?error="))
		s.Equal("test-state", response.Query["state"])
		s.False(response.HasKeyInQuery("code"))
		s.False(response.HasKeyInQuery("scope"))
		s.True(response.HasKeyInQuery("error"))
	})
}
