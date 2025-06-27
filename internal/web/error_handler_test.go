package web

import (
	"github.com/earaujoassis/space/test/utils"
)

func (s *WebHandlerTestSuite) TestErrorHandler() {
	w := s.PerformRequest(s.Router, "GET", "/error", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(200, w.Code)
	s.Contains(r.Body, "Unexpected Error")
}
