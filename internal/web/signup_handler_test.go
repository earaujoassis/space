package web

import (
	"github.com/earaujoassis/space/test/utils"
)

func (s *WebHandlerTestSuite) TestSignupHandler() {
	w := s.PerformRequest(s.Router, "GET", "/signup", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(200, w.Code)
	s.Contains(r.Body, "Sign Up")
	s.Contains(r.Body, "<script src=\"/public/js/io.min.js\"></script>")
}
