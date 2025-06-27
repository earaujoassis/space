package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestClientsProfileHandlerWithoutHeader() {
	client := s.Factory.NewClient().Model
	path := fmt.Sprintf("/api/clients/%s/profile", client.UUID)
	w := s.PerformRequest(s.Router, "PATCH", path, nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestClientsProfileHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	client := s.Factory.NewClient().Model
	path := fmt.Sprintf("/api/clients/%s/profile", client.UUID)

	w := s.PerformRequest(s.Router, "PATCH", path, header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *ApiHandlerTestSuite) TestClientsProfileHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	client := s.Factory.NewClient().Model
	path := fmt.Sprintf("/api/clients/%s/profile", client.UUID)
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)

	w := s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token string", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestClientsProfileHandlerInvalidId() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	path := "/api/clients/1/profile"
	formData := url.Values{}
	formData.Set("canonical_uri", "http://localhost:3000")
	formData.Set("redirect_uri", "http://localhost:3000/callback")
	formData.Set("scopes", "openid profile")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", path, header, cookie, strings.NewReader(encoded))
	s.Require().Equal(400, w.Code)

	path = "/api/clients/4862e6b00d95436d92b1b99eae84be8e/profile"
	formData = url.Values{}
	formData.Set("canonical_uri", "http://localhost:3000")
	formData.Set("redirect_uri", "http://localhost:3000/callback")
	formData.Set("scopes", "openid profile")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", path, header, cookie, strings.NewReader(encoded))
	s.Require().Equal(400, w.Code)
}

func (s *ApiHandlerTestSuite) TestClientsProfileHandlerByAdminUser() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	client := s.Factory.NewClient().Model
	path := fmt.Sprintf("/api/clients/%s/profile", client.UUID)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "PATCH", path, header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])

	formData := url.Values{}
	formData.Set("canonical_uri", "http://localhost:4000")
	formData.Set("redirect_uri", "http://localhost:3000/callback")
	formData.Set("scopes", "openid profile unknown")
	encoded := formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", path, header, cookie, strings.NewReader(encoded))
	s.Require().Equal(400, w.Code)

	w = s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("canonical_uri", "http://localhost:3000")
	formData.Set("redirect_uri", "http://localhost:3000/callback")
	formData.Set("scopes", "openid profile")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", path, header, cookie, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}

func (s *ApiHandlerTestSuite) TestClientsProfileHandlerByCommonUser() {
	cookie := s.createSessionCookie(false)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	client := s.Factory.NewClient().Model
	path := fmt.Sprintf("/api/clients/%s/profile", client.UUID)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	formData := url.Values{}
	formData.Set("canonical_uri", "http://localhost:3000")
	formData.Set("redirect_uri", "http://localhost:3000/callback")
	formData.Set("scopes", "openid profile")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", path, header, cookie, strings.NewReader(encoded))
	s.Require().Equal(401, w.Code)
}
