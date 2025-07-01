package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestUsersCreateHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "POST", "/api/users", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestUsersCreateHandlerWhenFeatureIsDisabled() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	s.AppCtx.FeatureGate.Disable("user.create")
	w := s.PerformRequest(s.Router, "POST", "/api/users", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(403, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("feature is not available at this time", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersCreateHandlerWithoutData() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	s.AppCtx.FeatureGate.Enable("user.create")
	w := s.PerformRequest(s.Router, "POST", "/api/users", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("missing essential fields", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersCreateHandler() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	s.AppCtx.FeatureGate.Enable("user.create")
	formData := url.Values{}
	formData.Set("first_name", gofakeit.FirstName())
	formData.Set("last_name", gofakeit.LastName())
	formData.Set("username", gofakeit.Username())
	formData.Set("email", gofakeit.Email())
	formData.Set("password", gofakeit.Password(true, true, true, true, false, 32))
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "POST", "/api/users", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(200, w.Code)
	s.True(r.HasKeyInJSON("user"))
	s.True(r.HasKeyInJSON("recover_secret"))
	s.True(r.HasKeyInJSON("code_secret_image"))
}
