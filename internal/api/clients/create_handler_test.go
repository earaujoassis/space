package clients

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit/v7"

	"github.com/earaujoassis/space/test/utils"
)

func (s *ClientsTestSuite) TestCreateHandlerWithoutHeader() {
	w := s.PerformRequest(s.Router, "POST", "/api/clients", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.Contains(r.Body, "missing X-Requested-By header attribute or Origin header does not comply with the same-origin policy")
}

func (s *ClientsTestSuite) TestCreateHandlerByUnauthenticatedUser() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/clients", header, nil, nil)
	s.Require().Equal(401, w.Code)
}

func (s *ClientsTestSuite) TestCreateHandlerWithoutActionGrant() {
	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
	}

	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)

	w := s.PerformRequest(s.Router, "POST", "/api/clients", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.Equal("must use valid token field", r.JSON["error"])
}

func (s *ClientsTestSuite) TestCreateHandlerByAdminUser() {
	cookie := s.CreateSessionCookie(true)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/clients", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(400, w.Code)
	s.True(r.HasKeyInJSON("error"))
	s.True(r.HasKeyInJSON("_message"))
	s.Equal("Client was not created", r.JSON["_message"])

	formData := url.Values{}
	formData.Set("name", gofakeit.Company())
	formData.Set("description", gofakeit.ProductDescription())
	formData.Set("canonical_uri", "http://localhost:3000")
	formData.Set("redirect_uri", "http://localhost:4000/callback")
	encoded := formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/clients", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(400, w.Code)

	formData = url.Values{}
	formData.Set("name", gofakeit.Company())
	formData.Set("description", gofakeit.ProductDescription())
	formData.Set("canonical_uri", "http://localhost")
	formData.Set("redirect_uri", "http://localhost/callback")
	encoded = formData.Encode()
	w = s.PerformRequest(s.Router, "POST", "/api/clients", header, cookie, strings.NewReader(encoded))
	s.Require().Equal(204, w.Code)
}

func (s *ClientsTestSuite) TestCreateHandlerByCommonUser() {
	cookie := s.CreateSessionCookie(false)
	s.NotNil(cookie)
	user := s.Factory.GetAvailableUser()
	actionToken := s.Factory.NewAction(user).Model.Token
	s.Require().Equal(len(actionToken), 64)

	header := &http.Header{
		"X-Requested-By": []string{"SpaceApi"},
		"Authorization":  []string{fmt.Sprintf("Bearer %s", actionToken)},
	}

	w := s.PerformRequest(s.Router, "POST", "/api/clients", header, cookie, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Require().Equal(403, w.Code)
	s.True(r.HasKeyInJSON("error"))
}
