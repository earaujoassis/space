package users

import (
	"fmt"
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *UsersTestSuite) TestClientsListHandlerWithoutHeader() {
	uuid := s.Factory.NewUser().Model.UUID
	path := fmt.Sprintf("/api/users/%s/clients", uuid)
	w := s.PerformRequest(s.Router, "GET", path, nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *UsersTestSuite) TestClientsListHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	uuid := s.Factory.NewUser().Model.UUID
	path := fmt.Sprintf("/api/users/%s/clients", uuid)

	w := s.PerformRequest(s.Router, "GET", path, header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *UsersTestSuite) TestClientsListHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)

	uuid := s.Factory.GetAvailableUser().UUID
	path := fmt.Sprintf("/api/users/%s/clients", uuid)
	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *UsersTestSuite) TestClientsListHandlerByAnotherUser() {
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
	path := fmt.Sprintf("/api/users/%s/clients", uuid)
	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])
}

func (s *UsersTestSuite) TestClientsListHandlerInvalidId() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	path := "/api/users/1/clients"
	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = "/api/users/4862e6b00d95436d92b1b99eae84be8e/clients"
	w = s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)
}

func (s *UsersTestSuite) TestClientsListHandler() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	client1 := s.Factory.NewClient().Model
	_ = s.Factory.NewRefreshToken(user, client1)
	client2 := s.Factory.NewClient().Model
	_ = s.Factory.NewRefreshToken(user, client2)

	path := fmt.Sprintf("/api/users/%s/clients", user.UUID)
	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("clients"))
	s.Equal(len(r.JSON["clients"].([]interface{})), 2)
}
