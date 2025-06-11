package integration

import (
	"github.com/earaujoassis/space/test/factory"
)

func (s *OAuthProviderSuite) TestAuthorizeUnsupportedGrants() {
	user := factory.NewUser()
	client := factory.NewClient()

	s.Run("should have a valid cookie", func() {
		s.Client.StartSession(user)
		s.True(s.Client.HasSessionCookie(), "Session cookie should be set")
	})

	s.Run("should return error if requesting unsupported Implicit Grant", func() {
		response := s.Client.GetAuthorize("token", client.Key, "http://localhost/callback", "test-state")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "unsupported_response_type")
	})
}
