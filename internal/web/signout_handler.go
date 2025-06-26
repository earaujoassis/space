package web

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func signoutHandler(c *gin.Context) {
	session := sessions.Default(c)

	userPublicID := session.Get("user_public_id")
	if userPublicID != nil {
		session.Delete("user_public_id")
		session.Save()
	}

	c.Redirect(http.StatusFound, "/signin")
}
