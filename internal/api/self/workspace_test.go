package self

import (
	"net/http"

	"github.com/earaujoassis/space/test/utils"
)

func (s *SelfTestSuite) TestWorkspaceHandler() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}
	w := s.PerformRequest(s.Router, "GET", "/api/users/me/workspace", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
	w = s.PerformRequest(s.Router, "GET", "/api/users/me/workspace", header, nil, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.Contains(r.Body, "access_denied")
	cookie := s.CreateSessionCookie(false)
	s.NotNil(cookie)
	w = s.PerformRequest(s.Router, "GET", "/api/users/me/workspace", header, cookie, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("workspace"))
}
