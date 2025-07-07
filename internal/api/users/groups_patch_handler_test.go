package users

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/earaujoassis/space/test/utils"
)

func (s *UsersTestSuite) TestGroupsPatchHandlerWithoutHeader() {
	userUuid := s.Factory.NewUser().Model.UUID
	clientUuid := s.Factory.NewClient().Model.UUID
	path := fmt.Sprintf("/api/users/%s/clients/%s/groups", userUuid, clientUuid)
	w := s.PerformRequest(s.Router, "PATCH", path, nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *UsersTestSuite) TestGroupsPatchHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	userUuid := s.Factory.NewUser().Model.UUID
	clientUuid := s.Factory.NewClient().Model.UUID
	path := fmt.Sprintf("/api/users/%s/clients/%s/groups", userUuid, clientUuid)

	w := s.PerformRequest(s.Router, "PATCH", path, header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *UsersTestSuite) TestGroupsPatchHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)

	userUuid := s.Factory.GetAvailableUser().UUID
	clientUuid := s.Factory.NewClient().Model.UUID
	path := fmt.Sprintf("/api/users/%s/clients/%s/groups", userUuid, clientUuid)
	w := s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *UsersTestSuite) TestGroupsPatchHandlerInvalidId() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	userUuid := s.Factory.NewUser().Model.UUID
	clientUuid := s.Factory.NewClient().Model.UUID

	path := fmt.Sprintf("/api/users/1/clients/%s/groups", clientUuid)
	w := s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/4862e6b00d95436d92b1b99eae84be8e/clients/%s/groups", clientUuid)
	w = s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/%s/clients/1/groups", userUuid)
	w = s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/%s/clients/4862e6b00d95436d92b1b99eae84be8e/groups", userUuid)
	w = s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)
}

func (s *UsersTestSuite) TestGroupsPatchHandlerWithoutMatchBetweenActionAndAuthenticatedUser() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	anotherUser := s.Factory.NewUser().Model
	actionToken := s.Factory.NewAction(anotherUser).Model.Token
	user := s.Factory.NewUser().Model
	client := s.Factory.NewClient().Model
	_ = s.Factory.NewGroup(user, client)
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	path := fmt.Sprintf("/api/users/%s/clients/%s/groups", user.UUID, client.UUID)
	w := s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	s.Require().Equal(401, w.Code)
}

func (s *UsersTestSuite) TestGroupsPatchHandlerWithoutData() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	authenticatedUser := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(authenticatedUser).Model.Token
	user := s.Factory.NewUser().Model
	client := s.Factory.NewClient().Model
	_ = s.Factory.NewGroup(user, client)
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	path := fmt.Sprintf("/api/users/%s/clients/%s/groups", user.UUID, client.UUID)
	w := s.PerformRequest(s.Router, "PATCH", path, header, cookie, nil)
	s.Require().Equal(204, w.Code)
}

func (s *UsersTestSuite) TestGroupsPatchHandler() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	authenticatedUser := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(authenticatedUser).Model.Token
	user := s.Factory.NewUser().Model
	client := s.Factory.NewClient().Model
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	formData := url.Values{}
	formData.Add("tags", "testing1")
	formData.Add("tags", "testing2")
	encoded := formData.Encode()
	path := fmt.Sprintf("/api/users/%s/clients/%s/groups", user.UUID, client.UUID)
	w := s.PerformRequest(s.Router, "PATCH", path, header, cookie, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}
