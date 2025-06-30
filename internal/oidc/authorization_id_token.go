package oidc

import (
	"slices"
	"strings"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/repository"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

// ImplicitFlowIDToken returns an OIDC id_token grant, given the right details
func ImplicitFlowIDToken(data utils.H, repositories *repository.RepositoryManager) (utils.H, error) {
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
	issuer := data["issuer"].(string)

	nonce := models.Nonce{
		ClientKey: client.Key,
		Code:      "",
		Nonce:     data["nonce"].(string),
	}
	if nonce.Nonce != "" {
		if !nonce.IsValid() {
			return shared.InvalidRequestResult(state)
		}
		if ok := repositories.Nonces().Create(nonce); !ok {
			return shared.InvalidRequestResult(state)
		}
	}

	if data["scope"] != nil {
		scope = data["scope"].(string)
	}

	if !slices.Contains(client.RedirectURI, redirectURI) {
		return shared.InvalidRequestResult(state)
	}

	if scope != "" && !client.HasRequestedScopes(utils.Scopes(scope)) && !strings.Contains(scope, models.OpenIDScope) {
		return shared.InvalidScopeResult(state)
	}

	idToken := createIDToken(issuer, user.PublicID, client.Key, data["nonce"].(string))
	session := models.Session{
		User:      user,
		Client:    client,
		IP:        ip,
		UserAgent: userAgent,
		Scopes:    scope,
		TokenType: models.GrantToken,
		Token:     idToken,
	}
	repositories.Sessions().Create(&session)
	if idToken != "" && session.ID > 0 {
		return utils.H{
			"id_token": idToken,
			"state":    state,
		}, nil
	}

	return shared.ServerErrorResult(state)
}
