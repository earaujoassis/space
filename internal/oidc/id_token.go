package oidc

import (
	"slices"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/services"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

func ImplicitFlowIdToken(data utils.H) (utils.H, error) {
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

	if data["scope"] != nil {
		scope = data["scope"].(string)
	}

	if !slices.Contains(client.RedirectURI, redirectURI) {
		return shared.InvalidRequestResult(state)
	}

	if scope == "" || !client.HasRequestedScopes(strings.Split(scope, " ")) || !strings.Contains(scope, models.OpenIDScope) {
		return shared.InvalidScopeResult(state)
	}

	idToken := createIDToken(issuer, user.PublicID, client.Key)
	session := services.CreateSessionWithToken(user, client, ip, userAgent, scope, models.IdToken, idToken)
	if idToken != "" && session.ID > 0 {
		return utils.H{
			"id_token": idToken,
			"state": state,
		}, nil
	}

	return shared.ServerErrorResult(state)
}

func createIDToken(issuer, userPublicId, clientKey string) string {
	keyManager, err := initKeyManager()
	if err != nil || len(keyManager.Keys) == 0 {
		logs.Propagatef(logs.Error, "JWKS is not available: %s", err)
		return ""
	}

	key := keyManager.Keys[0]
	claims := jwt.MapClaims{
		"iss": issuer,
		"sub": userPublicId,
		"aud": clientKey,
		"exp": time.Now().Add(time.Hour).Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token.Header["kid"] = key.ID

	signedToken, err := token.SignedString(key.PrivateKey)
	if err != nil {
		logs.Propagatef(logs.Error, "Could not sign id_token: %s", err)
		return ""
	}

	return signedToken
}
