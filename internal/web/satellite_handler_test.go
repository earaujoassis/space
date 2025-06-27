package web

import (
	"github.com/earaujoassis/space/test/utils"
)

func (s *WebHandlerTestSuite) TestSatelliteHandler() {
	w := s.PerformRequest(s.Router, "GET", "/", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(302, w.Code)
	s.Equal(r.Location, "/signin")

	cookie := s.createSessionCookie()
	s.Require().NotNil(cookie)

	w = s.PerformRequest(s.Router, "GET", "/", nil, cookie, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Equal(200, w.Code)
	s.Contains(r.Body, "Mission Control")
	s.Contains(r.Body, "<script src=\"/public/js/himalia.min.js\"></script>")
}
