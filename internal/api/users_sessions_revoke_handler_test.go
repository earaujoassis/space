package api

import (
	"fmt"
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestUsersSessionsRevokeHandlerWithoutHeader() {
	user := s.Factory.NewUser().Model
	session := s.Factory.NewApplicationSession(user).Model
	path := fmt.Sprintf("/api/users/%s/sessions/%s/revoke", user.UUID, session.UUID)
	w := s.PerformRequest(s.Router, "DELETE", path, nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestUsersSessionsRevokeHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	user := s.Factory.NewUser().Model
	session := s.Factory.NewApplicationSession(user).Model
	path := fmt.Sprintf("/api/users/%s/sessions/%s/revoke", user.UUID, session.UUID)

	w := s.PerformRequest(s.Router, "DELETE", path, header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *ApiHandlerTestSuite) TestUsersSessionsRevokeHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)

	user := s.Factory.NewUser().Model
	session := s.Factory.NewApplicationSession(user).Model
	path := fmt.Sprintf("/api/users/%s/sessions/%s/revoke", user.UUID, session.UUID)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token string", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersSessionsRevokeHandlerByAnotherUser() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	anotherUser := s.Factory.NewUser().Model
	session := s.Factory.NewApplicationSession(user).Model
	path := fmt.Sprintf("/api/users/%s/sessions/%s/revoke", anotherUser.UUID, session.UUID)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersSessionsRevokeHandlerOfAnotherUser() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	anotherUser := s.Factory.NewUser().Model
	anotherSession := s.Factory.NewApplicationSession(anotherUser).Model
	path := fmt.Sprintf("/api/users/%s/sessions/%s/revoke", user.UUID, anotherSession.UUID)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersSessionsRevokeHandlerInvalidId() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	session := s.Factory.NewApplicationSession(user).Model
	path := fmt.Sprintf("/api/users/1/sessions/%s/revoke", session.UUID)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/4862e6b00d95436d92b1b99eae84be8e/sessions/%s/revoke", session.UUID)
	w = s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/%s/sessions/1/revoke", user.UUID)
	w = s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/%s/sessions/4862e6b00d95436d92b1b99eae84be8e/revoke", user.UUID)
	w = s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)
}

func (s *ApiHandlerTestSuite) TestUsersSessionsRevokeHandler() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	session := s.Factory.NewApplicationSession(user).Model
	path := fmt.Sprintf("/api/users/%s/sessions/%s/revoke", user.UUID, session.UUID)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(204, w.Code)
}
