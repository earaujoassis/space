package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestServicesCreateHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "POST", "/api/services", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestServicesCreateHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/services", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.Contains(r.Body, "User must be authenticated")
}

func (s *ApiHandlerTestSuite) TestServicesCreateHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)

	w := s.PerformRequest(s.Router, "POST", "/api/services", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestServicesCreateHandlerWithoutData() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/services", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.True(r.HasKeyInJSON("_message"))
	s.Equal("Service was not created", r.JSON["_message"])
}

func (s *ApiHandlerTestSuite) TestServicesCreateHandlerByAdminUser() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	formData := url.Values{}
	formData.Set("name", "") // empty name
	formData.Set("description", gofakeit.ProductDescription())
	formData.Set("canonical_uri", "http://localhost")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/services", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(400, w.Code)

	formData = url.Values{}
	formData.Set("name", gofakeit.Company())
	formData.Set("description", gofakeit.ProductDescription())
	formData.Set("canonical_uri", "http://localhost")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/services", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}

func (s *ApiHandlerTestSuite) TestServicesCreateHandlerByCommonUser() {
	cookie := s.createSessionCookie(false)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/services", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])

	formData := url.Values{}
	formData.Set("name", gofakeit.Company())
	formData.Set("description", gofakeit.ProductDescription())
	formData.Set("canonical_uri", "http://localhost")
	encoded := formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/services", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])
}
