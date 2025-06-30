package web

import (
	"github.com/earaujoassis/space/test/utils"
)

func (s *WebHandlerTestSuite) TestSigninHandler() {
	w := s.PerformRequest(s.Router, "GET", "/signin", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(200, w.Code)
	s.Contains(r.Body, "Sign In")
	s.Contains(r.Body, "<script src=\"/public/js/ganymede.min.js\"></script>")
}
