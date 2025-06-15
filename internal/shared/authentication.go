package shared

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/security"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

// The following Authorization method is used by OAuth clients only
func ClientBasicAuthorization(c *gin.Context) {
	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)

	if !security.ValidBase64(authorizationBasic) {
		c.JSON(http.StatusBadRequest, utils.H{
			"error": "must use valid Authorization string",
		})
		c.Abort()
		return
	}

	client := ClientAuthentication(authorizationBasic)
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
	authorizationBearer := strings.Replace(c.Request.Header.Get("Authorization"), "Bearer ", "", 1)

	if !security.ValidToken(authorizationBearer) {
		c.JSON(http.StatusBadRequest, utils.H{
			"error": "must use valid token string",
		})
		c.Abort()
		return
	}

	session := AccessAuthentication(authorizationBearer)
	if session.IsNewRecord() || !services.SessionGrantsReadAbility(session) {
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

// ClientAuthentication authenticates a client application, extracting the key-secret pair;
//
//	and returns a client entry/model, given the key-secret pair
func ClientAuthentication(authorizationHeader string) models.Client {
	key, secret := BasicAuthDecode(authorizationHeader)
	return services.ClientAuthentication(key, secret)
}

// AccessAuthentication obtains a Session entry (typed as an `Access Token`) through
//
//	its token string
func AccessAuthentication(token string) models.Session {
	return services.FindSessionByToken(token, models.AccessToken)
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
