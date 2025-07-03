package integration

func (s *OAuthProviderSuite) TestSessionCreation() {
	s.Run("should create user session", func() {
		s.Client.ClearSession()
		user := s.Factory.NewUser()
		code := user.GenerateCode()

		response := s.Client.LoginUser(user.Model.Email, user.Passphrase, code)

		s.Equal(200, response.StatusCode)
		s.Equal("created", response.JSON["_status"])
		s.Equal("/session", response.JSON["redirect_uri"])
		s.Equal("public", response.JSON["scope"])
		s.Equal("authorization_code", response.JSON["grant_type"])
	})
}
