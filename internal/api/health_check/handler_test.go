package health_check

import (
	"github.com/earaujoassis/space/test/utils"
)

func (s *HealthCheckTestSuite) TestHandler() {
	w := s.PerformRequest(s.Router, "GET", "/api/health-check", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.Contains(r.Body, "healthy")
}
