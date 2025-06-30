package web

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/test/factory"
	"github.com/earaujoassis/space/test/unit"
	"github.com/earaujoassis/space/test/utils"
)

type WebHandlerTestSuite struct {
	unit.BaseTestSuite
	Router *gin.Engine
}

func (s *WebHandlerTestSuite) SetupSuite() {
	s.BaseTestSuite.SetupSuite()
	gin.SetMode(gin.TestMode)
	s.Router = s.setupRouter()
}

func (s *WebHandlerTestSuite) SetupTest() {
	s.BaseTestSuite.SetupTest()
	s.Router = s.setupRouter()
}

func (s *WebHandlerTestSuite) setupRouter() *gin.Engine {
	router := gin.New()
	store := cookie.NewStore([]byte(s.Config.SessionSecret))
	store.Options(sessions.Options{Secure: false, HttpOnly: true})
	router.Use(sessions.Sessions("space.session", store))
	router.Use(ioc.InjectAppContext(s.BaseTestSuite.AppCtx))
	router.GET("/set-session", func(c *gin.Context) {
		session := sessions.Default(c)
		repositories := ioc.GetRepositories(c)
		user := s.Factory.NewUserWithOption(factory.UserOptions{Admin: false}).Model
		client := repositories.Clients().FindOrCreate(models.DefaultClient)
		applicationSession := models.Session{
			User:      user,
			Client:    client,
			IP:        gofakeit.IPv4Address(),
			UserAgent: gofakeit.UserAgent(),
			Scopes:    models.PublicScope,
			TokenType: models.ApplicationToken,
		}
		err := repositories.Sessions().Create(&applicationSession)
		s.Require().NoError(err)
		session.Set(shared.CookieSessionKey, applicationSession.Token)
		session.Save()
		c.String(200, "Session set")
	})
	router.RedirectTrailingSlash = false
	ExposeRoutes(router)
	return router
}

func (s *WebHandlerTestSuite) createSessionCookie() *http.Cookie {
	w := s.PerformRequest(s.Router, "GET", "/set-session", nil, nil, nil)
	r := utils.ParseResponse(w.Result(), nil)
	s.Equal(200, w.Code)
	s.Contains(r.Body, "Session set")

	for _, cookie := range w.Result().Cookies() {
		if cookie.Name == "space.session" {
			return cookie
		}
	}

	return nil
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, new(WebHandlerTestSuite))
}
