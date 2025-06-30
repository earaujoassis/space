package api

import (
	"fmt"
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestClientsCredentialsHandlerByAdminUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.createSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header = &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	client := s.Factory.NewClient().Model
	path := fmt.Sprintf("/api/clients/%s/credentials", client.UUID)

	w := s.PerformRequest(s.Router, "GET", path, nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])

	header = &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w = s.PerformRequest(s.Router, "GET", path, header, nil, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("User must be authenticated", r.JSON["_message"])

	w = s.PerformRequest(s.Router, "GET", path, nil, cookie, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.Contains(r.Body, "name,client_key,client_secret")

	path = "/api/clients/1/credentials"
	w = s.PerformRequest(s.Router, "GET", path, nil, cookie, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("Client credentials are not available", r.JSON["_message"])

	path = "/api/clients/4862e6b00d95436d92b1b99eae84be8e/credentials"
	w = s.PerformRequest(s.Router, "GET", path, nil, cookie, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("Client credentials are not available", r.JSON["_message"])
}

func (s *ApiHandlerTestSuite) TestClientsCredentialsHandlerByCommonUser() {
	cookie := s.createSessionCookie(false)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	client := s.Factory.NewClient().Model
	path := fmt.Sprintf("/api/clients/%s/credentials", client.UUID)

	w := s.PerformRequest(s.Router, "GET", path, header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("access_denied", r.JSON["error"])
}
