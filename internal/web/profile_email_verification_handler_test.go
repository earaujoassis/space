package web

import (
	"fmt"

	"github.com/earaujoassis/space/test/utils"
)

func (s *WebHandlerTestSuite) TestProfileEmailVerificationHandler() {
	user := s.Factory.NewUser().Model
	actionSession := s.Factory.NewAction(user).Model
	s.Require().NotEmpty(actionSession.Token)

	w := s.PerformRequest(s.Router, "GET", "/profile/email_verification", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(401, w.Code)
	s.Contains(r.Body, "Email Confirmation")

	path := fmt.Sprintf("/profile/email_verification?_=%s", actionSession.Token)
	w = s.PerformRequest(s.Router, "GET", path, nil, nil, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Equal(302, w.Code)
	s.Equal("/", r.Location)
}
