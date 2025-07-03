package api

import (
	"fmt"
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestClientsListHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "GET", "/api/clients", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestClientsListHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "GET", "/api/clients", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])
}

func (s *ApiHandlerTestSuite) TestClientsListHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)

	w := s.PerformRequest(s.Router, "GET", "/api/clients", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestClientsListHandlerByAdminUser() {
	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)
	_ = s.Factory.NewClient().Model

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "GET", "/api/clients", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("clients"))
	clients := r.JSON["clients"].([]interface{})
	s.Equal(1, len(clients))
	client := clients[0].(map[string]interface{})
	s.NotEmpty(client["id"])
}

func (s *ApiHandlerTestSuite) TestClientsListHandlerByCommonUser() {
	cookie := s.createSessionCookie(false)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)
	_ = s.Factory.NewClient().Model

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "GET", "/api/clients", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])
}
