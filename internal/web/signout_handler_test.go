package web

import (
	"github.com/earaujoassis/space/test/utils"
)

func (s *WebHandlerTestSuite) TestSignoutHandler() {
	w := s.PerformRequest(s.Router, "GET", "/signout", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(302, w.Code)
	s.Equal(r.Location, "/signin")

	cookie := s.createSessionCookie()
	s.Require().NotNil(cookie)

	w = s.PerformRequest(s.Router, "GET", "/signout", nil, cookie, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Equal(302, w.Code)
	s.Equal(r.Location, "/signin")
}
