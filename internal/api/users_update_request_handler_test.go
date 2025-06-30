package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestUsersUpdateRequestHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "POST", "/api/users/update/request", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestUsersUpdateRequestHandlerWithoutHolder() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid holder string", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersUpdateRequestHandlerWithInvalidRequest() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("holder", gofakeit.Email())
	formData.Set("request_type", "invalid")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("request type not available", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersUpdateRequestHandlerWithUnknownHolder() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("holder", gofakeit.Email())
	formData.Set("request_type", "password")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("holder", gofakeit.Email())
	formData.Set("request_type", "secrets")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("holder", gofakeit.Email())
	formData.Set("request_type", "email_verification")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}

func (s *ApiHandlerTestSuite) TestUsersUpdateRequestHandlerWithKnownHolder() {
	userTest := s.Factory.NewUser()
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("holder", userTest.Email)
	formData.Set("request_type", "password")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("holder", userTest.Email)
	formData.Set("request_type", "secrets")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("holder", userTest.Email)
	formData.Set("request_type", "email_verification")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("holder", userTest.Username)
	formData.Set("request_type", "password")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("holder", userTest.Username)
	formData.Set("request_type", "secrets")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("holder", userTest.Username)
	formData.Set("request_type", "email_verification")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/users/update/request", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}
