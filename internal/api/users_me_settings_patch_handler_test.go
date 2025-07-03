package api

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestUsersMeSettingsPatchHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestUsersMeSettingsPatchHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *ApiHandlerTestSuite) TestUsersMeSettingsPatchHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)

	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMeSettingsPatchHandlerWithoutKey() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting key", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMeSettingsPatchHandlerWithInvalidKey() {
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
	formData.Set("key", "invalid")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting key", r.JSON["error"])

	formData = url.Values{}
	formData.Set("key", "invalid.three.parts")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMeSettingsPatchHandlerWithInvalidValue() {
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
	formData.Set("key", "notifications.system-email-notifications.authentication")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting", r.JSON["error"])

	formData = url.Values{}
	formData.Set("key", "notifications.system-email-notifications.authentication")
	formData.Set("value", "10")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMeSettingsPatchHandlerWithInvalidEmail() {
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
	formData.Set("key", "notifications.system-email-notifications.email-address")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting", r.JSON["error"])

	formData = url.Values{}
	formData.Set("key", "notifications.system-email-notifications.email-address")
	formData.Set("value", "example@example.com")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting", r.JSON["error"])

	_ = s.Factory.NewEmailFor(user).Model

	formData = url.Values{}
	formData.Set("key", "notifications.system-email-notifications.email-address")
	formData.Set("value", "example@example.com")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting", r.JSON["error"])

	formData = url.Values{}
	formData.Set("key", "notifications.system-email-notifications.email-address")
	formData.Set("value", "10")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Require().True(r.HasKeyInJSON("error"))
	s.Equal("invalid setting", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMeSettingsPatchHandler() {
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
	formData.Set("key", "notifications.system-email-notifications.email-address")
	formData.Set("value", user.Email)
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	email := s.Factory.NewEmailFor(user).Model

	formData = url.Values{}
	formData.Set("key", "notifications.system-email-notifications.email-address")
	formData.Set("value", email.Address)
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)

	formData = url.Values{}
	formData.Set("key", "notifications.system-email-notifications.authentication")
	formData.Set("value", "true")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/settings", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}
