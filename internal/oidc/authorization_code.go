package oidc

import (
	"strings"

	"golang.org/x/exp/slices"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

// AuthorizationCodeGrant returns an OIDC authorization code grant, given the right details
func AuthorizationCodeGrant(data utils.H) (utils.H, error) {
	var redirectURI string
	var scope string
	var state string

	var ip string
	var userAgent string

	var user models.User
	var client models.Client

	if data["redirect_uri"] == nil || data["user"] == nil || data["client"] == nil {
		return shared.InvalidRequestResult(state)
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

	nonce := data["nonce"].(string)
	if nonce != "" {
		if !isValidNonce(nonce) {
			return shared.InvalidRequestResult(state)
		}
		if ok := storeNonceForClient(client.Key, nonce); !ok {
			return shared.InvalidRequestResult(state)
		}
	}

	if data["scope"] != nil {
		scope = data["scope"].(string)
	}

	if !slices.Contains(client.RedirectURI, redirectURI) {
		return shared.InvalidRequestResult(state)
	}

	if scope == "" || !client.HasRequestedScopes(strings.Split(scope, " ")) || !strings.Contains(scope, models.OpenIDScope) {
		return shared.InvalidScopeResult(state)
	}

	session := services.CreateSession(user, client, ip, userAgent, scope, models.GrantToken)
	if session.ID > 0 {
		return utils.H{
			"code":  session.Token,
			"state": state,
			"scope": scope,
		}, nil
	}

	return shared.ServerErrorResult(state)
}
