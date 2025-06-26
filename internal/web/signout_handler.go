package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/shared"
)

func signoutHandler(c *gin.Context) {
	session := sessions.Default(c)

	sessionToken := session.Get(shared.CookieSessionKey)
	if sessionToken != nil {
		session.Delete(shared.CookieSessionKey)
		session.Save()
	}

	c.Redirect(http.StatusFound, "/signin")
}
