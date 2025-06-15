package oauth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func tokenHandler(c *gin.Context) {
	var grantType = c.PostForm("grant_type")

	authorizationBasic := strings.Replace(c.Request.Header.Get("Authorization"), "Basic ", "", 1)
	client := shared.ClientAuthentication(authorizationBasic)
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
		result, err := AccessTokenRequest(utils.H{
			"grant_type":   grantType,
			"code":         c.PostForm("code"),
			"redirect_uri": c.PostForm("redirect_uri"),
			"client":       client,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.H{
				"error": result["error"],
			})
			return
		}
		c.JSON(http.StatusOK, utils.H{
			"user_id":       result["user_id"],
			"access_token":  result["access_token"],
			"token_type":    result["token_type"],
			"expires_in":    result["expires_in"],
			"refresh_token": result["refresh_token"],
			"scope":         result["scope"],
		})
		return
	// Refreshing an Access Token
	case shared.RefreshToken:
		result, err := RefreshTokenRequest(utils.H{
			"grant_type":    grantType,
			"refresh_token": c.PostForm("refresh_token"),
			"scope":         c.PostForm("scope"),
			"client":        client,
		})
		if err != nil {
			c.JSON(http.StatusBadRequest, utils.H{
				"error": result["error"],
			})
			return
		}
		c.JSON(http.StatusOK, utils.H{
			"user_id":       result["user_id"],
			"access_token":  result["access_token"],
			"token_type":    result["token_type"],
			"expires_in":    result["expires_in"],
			"refresh_token": result["refresh_token"],
			"scope":         result["scope"],
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
