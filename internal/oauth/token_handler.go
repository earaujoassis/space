package oauth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/ioc"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func tokenHandler(c *gin.Context) {
	var grantType = c.PostForm("grant_type")

	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)
	key, secret := shared.BasicAuthDecode(authorizationBasic)
	repositories := ioc.GetRepositories(c)
	client := repositories.Clients().Authentication(key, secret)
	if client.IsNewRecord() {
		c.Header("WWW-Authenticate", "Basic realm=\"OAuth\"")
		c.JSON(http.StatusUnauthorized, utils.H{
			"error":             shared.InvalidClient,
			"error_description": "Client authentication failed",
		})
		return
	}

	switch grantType {
	// Authorization Code Grant
	case shared.AuthorizationCode:
		result, err := AccessTokenRequest(AccessTokenParams{
			GrantType:   grantType,
			Code:        c.PostForm("code"),
			RedirectURI: c.PostForm("redirect_uri"),
			Client:      client,
		}, repositories)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.H{
				"error": err.ErrorType,
			})
			return
		}
		c.JSON(http.StatusOK, utils.H{
			"access_token":  result.AccessToken,
			"token_type":    result.TokenType,
			"expires_in":    result.ExpiresIn,
			"refresh_token": result.RefreshToken,
		})
		return
	// Refreshing an Access Token
	case shared.RefreshToken:
		result, err := RefreshTokenRequest(RefreshTokenParams{
			GrantType:    grantType,
			RefreshToken: c.PostForm("refresh_token"),
			Scope:        c.PostForm("scope"),
			Client:       client,
		}, repositories)
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.H{
				"error": err.ErrorType,
			})
			return
		}
		c.JSON(http.StatusOK, utils.H{
			"access_token":  result.AccessToken,
			"token_type":    result.TokenType,
			"expires_in":    result.ExpiresIn,
			"refresh_token": result.RefreshToken,
		})
		return
	// Resource Owner Password Credentials Grant
	// Client Credentials Grant
	case shared.Password, shared.ClientCredentials:
		c.JSON(http.StatusBadRequest, utils.H{
			"error": shared.UnsupportedGrantType,
		})
		return
	default:
		c.JSON(http.StatusBadRequest, utils.H{
			"error": shared.InvalidRequest,
		})
		return
	}
}
