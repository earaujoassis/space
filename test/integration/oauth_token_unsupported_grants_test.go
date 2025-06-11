package integration

import (
	"github.com/earaujoassis/space/test/factory"
)

func (s *OAuthProviderSuite) TestTokenUnsupportedGrants() {
	// user := factory.NewUser()
	client := factory.NewClient()

	s.Run("should return error if requesting unsupported Resource Owner Password Credentials Grant", func() {
		response := s.Client.PostToken(client.BasicAuthEncode(), "password")

		s.Equal(405, response.StatusCode)
		s.Contains(response.Body, "unsupported_grant_type")
		json := response.JSON
		s.Equal(json["_status"], "error")
		s.Equal(json["error"], "unsupported_grant_type")
	})

	s.Run("should return error if requesting unsupported Client Credentials Grant", func() {
		response := s.Client.PostToken(client.BasicAuthEncode(), "client_credentials")

		s.Equal(405, response.StatusCode)
		s.Contains(response.Body, "unsupported_grant_type")
		json := response.JSON
		s.Equal(json["_status"], "error")
		s.Equal(json["error"], "unsupported_grant_type")
	})
}
