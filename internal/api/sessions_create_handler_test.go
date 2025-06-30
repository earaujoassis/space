package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestSessionsCreateHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/create", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestSessionsCreateHandlerWithoutData() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/sessions/create", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid holder string", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestSessionsCreateHandler() {
	userTest := s.Factory.NewUser()

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("holder", userTest.Username)
	formData.Set("password", userTest.Passphrase)
	formData.Set("passcode", userTest.GenerateCode())
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/create", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("scope"))
	s.True(r.HasKeyInJSON("grant_type"))
	s.True(r.HasKeyInJSON("code"))
	s.True(r.HasKeyInJSON("redirect_uri"))
	s.True(r.HasKeyInJSON("client_id"))
	s.True(r.HasKeyInJSON("state"))

	formData = url.Values{}
	formData.Set("holder", userTest.Email)
	formData.Set("password", userTest.Passphrase)
	formData.Set("passcode", userTest.GenerateCode())
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/sessions/create", header, nil, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("scope"))
	s.True(r.HasKeyInJSON("grant_type"))
	s.True(r.HasKeyInJSON("code"))
	s.True(r.HasKeyInJSON("redirect_uri"))
	s.True(r.HasKeyInJSON("client_id"))
	s.True(r.HasKeyInJSON("state"))

	formData = url.Values{}
	formData.Set("holder", "anotherimprobableemail@example.com")
	formData.Set("password", userTest.Passphrase)
	formData.Set("passcode", userTest.GenerateCode())
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/sessions/create", header, nil, strings.NewReader(encoded))
	s.Require().Equal(400, w.Code)

	formData = url.Values{}
	formData.Set("holder", userTest.Email)
	formData.Set("password", userTest.Passphrase)
	formData.Set("passcode", "000000")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/sessions/create", header, nil, strings.NewReader(encoded))
	s.Require().Equal(400, w.Code)
}
