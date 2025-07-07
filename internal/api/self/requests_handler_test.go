package self

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/test/utils"
)

func (s *SelfTestSuite) TestRequestsHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "POST", "/api/users/me/requests", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *SelfTestSuite) TestRequestsHandlerWithoutRequestType() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("request type not available", r.JSON["error"])
}

func (s *SelfTestSuite) TestRequestsHandlerWithInvalidRequest() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("request_type", "invalid")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("request type not available", r.JSON["error"])
}

func (s *SelfTestSuite) TestRequestsHandlerWithUnknownHolder() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("request_type", "password")
	formData.Set("holder", gofakeit.Email())
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid holder field", r.JSON["error"])

	formData = url.Values{}
	formData.Set("request_type", "secrets")
	formData.Set("holder", gofakeit.Email())
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid holder field", r.JSON["error"])

	email := gofakeit.Email()
	formData = url.Values{}
	formData.Set("request_type", "email_verification")
	formData.Set("holder", email)
	formData.Set("email", email)
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid holder field", r.JSON["error"])
}

func (s *SelfTestSuite) TestRequestsHandlerWithKnownHolder() {
	userTest := s.Factory.NewUser()
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("request_type", "password")
	formData.Set("holder", userTest.Model.Email)
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("request_type", "secrets")
	formData.Set("holder", userTest.Model.Email)
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("request_type", "email_verification")
	formData.Set("holder", userTest.Model.Email)
	formData.Set("email", userTest.Model.Email)
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}

func (s *SelfTestSuite) TestRequestsHandlerWithoutEmailForValidation() {
	userTest := s.Factory.NewUser()
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("request_type", "email_verification")
	formData.Set("holder", userTest.Model.Username)
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/users/me/requests", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid email field", r.JSON["error"])
}
