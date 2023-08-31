package oauth

import (
	"strings"

	"golang.org/x/exp/slices"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/utils"
)

// AuthorizationCodeGrant returns an OAuth 2 authorization code grant, given the right details
func AuthorizationCodeGrant(data utils.H) (utils.H, error) {
	var redirectURI string
	var scope string
	var state string

	var ip string
	var userAgent string

	var user models.User
	var client models.Client

	if data["redirect_uri"] == nil || data["user"] == nil || data["client"] == nil {
		return invalidRequestResult(state)
	}

	if data["state"] != nil {
		state = data["state"].(string)
	}

	if data["ip"] != nil {
		ip = data["ip"].(string)
	}

	if data["userAgent"] != nil {
		userAgent = data["userAgent"].(string)
	}

	redirectURI = data["redirect_uri"].(string)
	client = data["client"].(models.Client)
	user = data["user"].(models.User)

	if data["scope"] != nil {
		scope = data["scope"].(string)
	}

	if !slices.Contains(client.RedirectURI, redirectURI) {
		return invalidRedirectURIResult(state)
	}

	/*
	 * WARNING
	 * It will grant access, but with a public-only scope
	 */
	if scope != "" && !strings.Contains(client.Scopes, scope) {
		scope = models.PublicScope
	}

	session := services.CreateSession(user, client, ip, userAgent, scope, models.GrantToken)
	if session.ID > 0 {
		return utils.H{
			"code":  session.Token,
			"state": state,
			"scope": scope,
		}, nil
	}

	return serverErrorResult(state)
}
