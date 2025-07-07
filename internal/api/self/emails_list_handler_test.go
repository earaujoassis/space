package self

import (
	"fmt"
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *SelfTestSuite) TestEmailsListHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "GET", "/api/users/me/emails", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *SelfTestSuite) TestEmailsListHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "GET", "/api/users/me/emails", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
}

func (s *SelfTestSuite) TestEmailsListHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)

	w := s.PerformRequest(s.Router, "GET", "/api/users/me/emails", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *SelfTestSuite) TestEmailsListHandler() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	_ = s.Factory.NewEmailFor(user)
	_ = s.Factory.NewEmailFor(user)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "GET", "/api/users/me/emails", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.Require().True(r.HasKeyInJSON("emails"))
	emails := r.JSON["emails"].([]interface{})
	s.Equal(2, len(emails))
	email := emails[0].(map[string]interface{})
	s.NotEmpty(email["id"])
	s.NotEmpty(email["address"])
	s.NotNil(email["verified"])
}
