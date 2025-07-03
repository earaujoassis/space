package api

import (
	"fmt"
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestUsersClientsRevokeHandlerWithoutHeader() {
	userUuid := s.Factory.NewUser().Model.UUID
	clientUuid := s.Factory.NewClient().Model.UUID
	path := fmt.Sprintf("/api/users/%s/clients/%s/revoke", userUuid, clientUuid)
	w := s.PerformRequest(s.Router, "DELETE", path, nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestUsersClientsRevokeHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	userUuid := s.Factory.NewUser().Model.UUID
	clientUuid := s.Factory.NewClient().Model.UUID
	path := fmt.Sprintf("/api/users/%s/clients/%s/revoke", userUuid, clientUuid)

	w := s.PerformRequest(s.Router, "DELETE", path, header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *ApiHandlerTestSuite) TestUsersClientsRevokeHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)

	userUuid := s.Factory.GetAvailableUser().UUID
	clientUuid := s.Factory.NewClient().Model.UUID
	path := fmt.Sprintf("/api/users/%s/clients/%s/revoke", userUuid, clientUuid)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersClientsRevokeHandlerByAnotherUser() {
	cookie := s.createSessionCookie(true)
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
	path := fmt.Sprintf("/api/users/%s/clients/%s/revoke", userUuid, clientUuid)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersClientsRevokeHandlerInvalidId() {
	cookie := s.createSessionCookie(true)
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

	path := fmt.Sprintf("/api/users/1/clients/%s/revoke", clientUuid)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/4862e6b00d95436d92b1b99eae84be8e/clients/%s/revoke", clientUuid)
	w = s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/%s/clients/1/revoke", userUuid)
	w = s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)

	path = fmt.Sprintf("/api/users/%s/clients/4862e6b00d95436d92b1b99eae84be8e/revoke", userUuid)
	w = s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(400, w.Code)
}

func (s *ApiHandlerTestSuite) TestUsersClientsRevokeHandler() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	client := s.Factory.NewClient().Model
	_ = s.Factory.NewRefreshToken(user, client)

	path := fmt.Sprintf("/api/users/%s/clients/%s/revoke", user.UUID, client.UUID)
	w := s.PerformRequest(s.Router, "DELETE", path, header, cookie, nil)
	s.Require().Equal(204, w.Code)
}
