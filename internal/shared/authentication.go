package shared

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/utils"
)

// The following Authorization method is used by OAuth clients only
func ClientBasicAuthorization(c *gin.Context) {
	repositories := ioc.GetRepositories(c)

	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)
	if !security.ValidBase64(authorizationBasic) {
		c.JSON(http.StatusBadRequest, utils.H{
			"error": "must use valid Authorization string",
		})
		c.Abort()
		return
	}

	key, secret := BasicAuthDecode(authorizationBasic)
	client := repositories.Clients().Authentication(key, secret)
	if client.IsNewRecord() {
		c.Header("WWW-Authenticate", "Basic realm=\"OAuth\"")
		c.JSON(http.StatusUnauthorized, utils.H{
			"error": AccessDenied,
		})
		c.Abort()
		return
	}
	c.Set("Client", client)
	c.Next()
}

// The following Authorization method is used by the OAuth clients, with an OAuth session token
func OAuthTokenBearerAuthorization(c *gin.Context) {
	repositories := ioc.GetRepositories(c)

	authorizationBearer := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)
	if !security.ValidToken(authorizationBearer) {
		c.JSON(http.StatusBadRequest, utils.H{
			"error": "must use valid token string",
		})
		c.Abort()
		return
	}

	session := repositories.Sessions().FindByToken(authorizationBearer, models.AccessToken)
	if session.IsNewRecord() || !session.GrantsReadAbility() {
		c.Header("WWW-Authenticate", "Basic realm=\"OAuth\"")
		c.JSON(http.StatusUnauthorized, utils.H{
			"error": AccessDenied,
		})
		c.Abort()
		return
	}
	c.Set("Session", session)
	c.Next()
}

func Scheme(request *http.Request) string {
	if scheme := request.Header.Get("X-Forwarded-Proto"); scheme != "" {
		return scheme
	}
	if request.TLS == nil {
		return "http"
	}

	return "https"
}
