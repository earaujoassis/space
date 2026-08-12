package sessions

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/test/utils"
)

func (s *SessionsTestSuite) TestRequestsHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/requests", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *SessionsTestSuite) TestRequestsHandlerWithoutRequestType() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/sessions/requests", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("request type not available", r.JSON["error"])
}

func (s *SessionsTestSuite) TestRequestsHandlerWithoutHolder() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("request_type", "passwordless_signin")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/requests", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid holder field", r.JSON["error"])
}

func (s *SessionsTestSuite) TestRequestsHandlerWithInvalidRequest() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("request_type", "invalid")
	formData.Set("holder", gofakeit.Email())
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/requests", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("request type not available", r.JSON["error"])
}

func (s *SessionsTestSuite) TestRequestsHandlerWithUnknownHolder() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("request_type", "passwordless_signin")
	formData.Set("holder", "anotherimprobableemail@example.com")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/requests", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}

func (s *SessionsTestSuite) TestRequestsHandler() {
	userTest := s.Factory.NewUser()

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("request_type", "passwordless_signin")
	formData.Set("holder", userTest.Model.Username)
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/sessions/requests", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}
