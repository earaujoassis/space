package web

import (
	"fmt"

	"github.com/earaujoassis/space/test/utils"
)

func (s *WebHandlerTestSuite) TestSessionHandler() {
	user := s.Factory.NewUser().Model
	grantSession := s.Factory.NewGrantToken(user).Model
	client := s.Factory.DefaultClient().Model

	w := s.PerformRequest(s.Router, "GET", "/session", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(302, w.Code)
	s.Equal("/signin", r.Location)

	path := fmt.Sprintf("/session?client_id=%s&code=%s&grant_type=authorization_code&scope=public&state=", client.Key, grantSession.Token)
	w = s.PerformRequest(s.Router, "GET", path, nil, nil, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Equal(302, w.Code)
	s.Equal("/", r.Location)

	grantSession = s.Factory.NewGrantToken(user).Model
	path = fmt.Sprintf("/session?client_id=%s&code=%s&grant_type=authorization_code&scope=public&state=&_=%%2Fprofile", client.Key, grantSession.Token)
	w = s.PerformRequest(s.Router, "GET", path, nil, nil, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Equal(302, w.Code)
	s.Equal("/profile", r.Location)

	cookie := s.createSessionCookie()
	s.Require().NotNil(cookie)
	w = s.PerformRequest(s.Router, "GET", "/session", nil, cookie, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Equal(302, w.Code)
	s.Equal("/", r.Location)
}
