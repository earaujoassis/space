package users

import (
	"fmt"
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *UsersTestSuite) TestSessionsListHandlerWithoutHeader() {
	uuid := s.Factory.NewUser().Model.UUID
	path := fmt.Sprintf("/api/users/%s/sessions", uuid)
	w := s.PerformRequest(s.Router, "GET", path, nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *UsersTestSuite) TestSessionsListHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	uuid := s.Factory.NewUser().Model.UUID
	path := fmt.Sprintf("/api/users/%s/sessions", uuid)

	w := s.PerformRequest(s.Router, "GET", path, header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *UsersTestSuite) TestSessionsListHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)

	uuid := s.Factory.GetAvailableUser().UUID
	path := fmt.Sprintf("/api/users/%s/sessions", uuid)
	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *UsersTestSuite) TestSessionsListHandlerByAnotherUser() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	uuid := s.Factory.NewUser().Model.UUID
	path := fmt.Sprintf("/api/users/%s/sessions", uuid)
	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])
}

func (s *UsersTestSuite) TestSessionsListHandlerInvalidId() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	path := "/api/users/1/sessions"
	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = "/api/users/4862e6b00d95436d92b1b99eae84be8e/sessions"
	w = s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)
}

func (s *UsersTestSuite) TestSessionsListHandler() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	_ = s.Factory.NewApplicationSession(user)
	_ = s.Factory.NewApplicationSession(user)

	uuid := user.UUID
	path := fmt.Sprintf("/api/users/%s/sessions", uuid)
	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("sessions"))
	s.Equal(len(r.JSON["sessions"].([]interface{})), 3) // 2 + 1 active
}
