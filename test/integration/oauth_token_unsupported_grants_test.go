package integration

import (
	"github.com/earaujoassis/space/test/factory"
)

func (s *OAuthProviderSuite) TestTokenUnsupportedGrants() {
	client := factory.NewClient()

	s.Run("should return error if requesting unsupported Resource Owner Password Credentials Grant", func() {
		response := s.Client.PostToken(client.BasicAuthEncode(), "password")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "unsupported_grant_type")
		json := response.JSON
		s.Equal("error", json["_status"], )
		s.Equal("unsupported_grant_type", json["error"])
	})

	s.Run("should return error if requesting unsupported Client Credentials Grant", func() {
		response := s.Client.PostToken(client.BasicAuthEncode(), "client_credentials")

		s.Equal(400, response.StatusCode)
		s.Contains(response.Body, "unsupported_grant_type")
		json := response.JSON
		s.Equal("error", json["_status"])
		s.Equal("unsupported_grant_type", json["error"])
	})
}
