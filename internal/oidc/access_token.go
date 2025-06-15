package oidc

import (
	"golang.org/x/exp/slices"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

// AccessTokenRequest returns OAuth 2 access and refresh tokens, given the right details:
//
//	basically, a `code` from `AuthorizationCodeGrant`
func AccessTokenRequest(data utils.H) (utils.H, error) {
	var user models.User
	var client models.Client

	var code string
	var redirectURI string

	if data["code"] == nil || data["redirect_uri"] == nil || data["client"] == nil {
		return shared.InvalidRequestResult("")
	}

	redirectURI = data["redirect_uri"].(string)
	code = data["code"].(string)
	client = data["client"].(models.Client)
	issuer := data["issuer"].(string)

	authorizationSession := services.FindSessionByToken(code, models.GrantToken)
	defer services.InvalidateSession(authorizationSession)
	if authorizationSession.IsNewRecord() {
		return shared.InvalidGrantResult("")
	}
	user = authorizationSession.User
	user = services.FindUserByPublicID(user.PublicID)
	if authorizationSession.Client.ID != client.ID {
		return shared.InvalidGrantResult("")
	}
	if !slices.Contains(authorizationSession.Client.RedirectURI, redirectURI) {
		return shared.InvalidGrantResult("")
	}

	accessToken := services.CreateSession(user,
		client,
		authorizationSession.IP,
		authorizationSession.UserAgent,
		authorizationSession.Scopes,
		models.AccessToken)
	refreshToken := services.CreateSession(user,
		client,
		authorizationSession.IP,
		authorizationSession.UserAgent,
		authorizationSession.Scopes,
		models.RefreshToken)

	if accessToken.IsNewRecord() || refreshToken.IsNewRecord() {
		return shared.ServerErrorResult("")
	}

	nonce := retrieveNonceForCode(code)
	idToken := createIDToken(issuer, user.PublicID, client.Key, nonce)
	return utils.H{
		"access_token":  accessToken.Token,
		"token_type":    "Bearer",
		"expires_in":    accessToken.ExpiresIn,
		"refresh_token": refreshToken.Token,
		"id_token":      idToken,
	}, nil
}
