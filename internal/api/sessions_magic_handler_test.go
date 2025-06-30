package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestSessionsMagicHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/magic", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestSessionsMagicHandlerWithoutHolder() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/sessions/magic", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid holder string", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestSessionsMagicHandlerWithUnknownHolder() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("holder", "anotherimprobableemail@example.com")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/magic", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}

func (s *ApiHandlerTestSuite) TestSessionsMagicHandler() {
	userTest := s.Factory.NewUser()

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("holder", userTest.Username)
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/magic", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}
