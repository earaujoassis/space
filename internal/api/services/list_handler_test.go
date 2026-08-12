package services

import (
	"fmt"
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ServicesTestSuite) TestListHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "GET", "/api/services", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ServicesTestSuite) TestListHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "GET", "/api/services", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
}

func (s *ServicesTestSuite) TestListHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)

	w := s.PerformRequest(s.Router, "GET", "/api/services", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *ServicesTestSuite) TestListHandlerByAdminUser() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	_ = s.Factory.NewService()
	_ = s.Factory.NewService()

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "GET", "/api/services", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("services"))
	services := r.JSON["services"].([]interface{})
	s.Equal(2, len(services))
	service := services[0].(map[string]interface{})
	s.NotEmpty(service["id"])
}

func (s *ServicesTestSuite) TestListHandlerByCommonUser() {
	cookie := s.CreateSessionCookie(false)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	_ = s.Factory.NewService()
	_ = s.Factory.NewService()

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "GET", "/api/services", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("services"))
	services := r.JSON["services"].([]interface{})
	s.Equal(2, len(services))
	service := services[0].(map[string]interface{})
	s.NotEmpty(service["id"])
}
