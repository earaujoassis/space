package self

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/earaujoassis/space/test/utils"
)

func (s *SelfTestSuite) TestAdminHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *SelfTestSuite) TestAdminHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *SelfTestSuite) TestAdminHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)

	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *SelfTestSuite) TestAdminHandlerWhenFeatureIsDisabled() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	s.AppCtx.FeatureGate.Disable("user.adminify")
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(403, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("feature is not available at this time", r.JSON["error"])
}

func (s *SelfTestSuite) TestAdminHandlerWithoutKey() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	s.AppCtx.FeatureGate.Enable("user.adminify")
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("application key is incorrect", r.JSON["error"])
}

func (s *SelfTestSuite) TestAdminHandlerWithoutCorrectKey() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	s.AppCtx.FeatureGate.Enable("user.adminify")
	formData := url.Values{}
	formData.Set("application_key", "incorrect")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("application key is incorrect", r.JSON["error"])
}

func (s *SelfTestSuite) TestAdminHandlerWithoutUserId() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	s.AppCtx.FeatureGate.Enable("user.adminify")
	formData := url.Values{}
	formData.Set("application_key", "masterapplicationkey")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid UUID for identification", r.JSON["error"])
}

func (s *SelfTestSuite) TestAdminHandlerWithIncorrectUserId() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	formData := url.Values{}
	formData.Set("user_id", "4862e6b00d95436d92b1b99eae84be8e")
	formData.Set("application_key", "masterapplicationkey")
	encoded := formData.Encode()
	s.AppCtx.FeatureGate.Enable("user.adminify")
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid UUID for identification", r.JSON["error"])

	s.AppCtx.FeatureGate.Enable("user.adminify")
	formData = url.Values{}
	formData.Set("user_id", "1")
	formData.Set("application_key", "masterapplicationkey")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid UUID for identification", r.JSON["error"])
}

func (s *SelfTestSuite) TestAdminHandlerByAnotherUser() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	anotherUser := s.Factory.NewUser().Model
	s.AppCtx.FeatureGate.Enable("user.adminify")
	formData := url.Values{}
	formData.Set("user_id", anotherUser.UUID)
	formData.Set("application_key", "masterapplicationkey")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(401, w.Code)
}

func (s *SelfTestSuite) TestAdminHandler() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	s.AppCtx.FeatureGate.Enable("user.adminify")
	formData := url.Values{}
	formData.Set("user_id", user.UUID)
	formData.Set("application_key", "masterapplicationkey")
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/admin", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}
