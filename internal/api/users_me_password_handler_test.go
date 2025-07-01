package api

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ApiHandlerTestSuite) TestUsersMePasswordHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/password", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ApiHandlerTestSuite) TestUsersMePasswordHandlerWithoutToken() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/password", header, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token string", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMePasswordHandlerWithoutData() {
	user := s.Factory.NewUser().Model
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("_", actionToken)
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/password", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("invalid password update attempt", r.JSON["error"])

	formData = url.Values{}
	formData.Set("_", actionToken)
	formData.Set("new_password", "")
	formData.Set("password_confirmation", "")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/password", header, nil, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("invalid password update attempt", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMePasswordHandlerIncorrectData() {
	user := s.Factory.NewUser().Model
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	formData := url.Values{}
	formData.Set("_", actionToken)
	formData.Set("new_password", gofakeit.Password(true, true, true, true, false, 12))
	formData.Set("password_confirmation", gofakeit.Password(true, true, true, true, false, 12))
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/password", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("new password and password confirmation must match each other", r.JSON["error"])

	password := gofakeit.Password(true, true, true, true, false, 9)
	formData = url.Values{}
	formData.Set("_", actionToken)
	formData.Set("new_password", password)
	formData.Set("password_confirmation", password)
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "PATCH", "/api/users/me/password", header, nil, strings.NewReader(encoded))
	r = utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("invalid password update attempt", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMePasswordHandlerWithoutPermission() {
	user := s.Factory.NewUser().Model
	actionToken := s.Factory.NewActionWithoutPermissions(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	password := gofakeit.Password(true, true, true, true, false, 10)
	formData := url.Values{}
	formData.Set("_", actionToken)
	formData.Set("new_password", password)
	formData.Set("password_confirmation", password)
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/password", header, nil, strings.NewReader(encoded))
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(401, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("invalid token string", r.JSON["error"])
}

func (s *ApiHandlerTestSuite) TestUsersMePasswordHandler() {
	user := s.Factory.NewUser().Model
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	password := gofakeit.Password(true, true, true, true, false, 10)
	formData := url.Values{}
	formData.Set("_", actionToken)
	formData.Set("new_password", password)
	formData.Set("password_confirmation", password)
	encoded := formData.Encode()
	w := s.PerformRequest(s.Router, "PATCH", "/api/users/me/password", header, nil, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}
