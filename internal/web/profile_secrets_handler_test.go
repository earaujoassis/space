package web

import (
	"fmt"

	"github.com/earaujoassis/space/test/utils"
)

func (s *WebHandlerTestSuite) TestProfileSecretsHandler() {
	user := s.Factory.NewUser().Model
	actionSession := s.Factory.NewAction(user).Model
	s.Require().NotEmpty(actionSession.Token)

	w := s.PerformRequest(s.Router, "GET", "/profile/secrets", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(401, w.Code)
	s.Contains(r.Body, "Update Resource Owner Credential")

	path := fmt.Sprintf("/profile/secrets?_=%s", actionSession.Token)
	w = s.PerformRequest(s.Router, "GET", path, nil, nil, nil)
	r = utils.ParseResponse(w.Result(), nil)
	s.Equal(200, w.Code)
	s.Contains(r.Body, "Update Resource Owner Credential")
	s.Contains(r.Body, "<h2>Yay!</h2>")
}
