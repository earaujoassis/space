package oauth

import (
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/repository"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

// RefreshTokenRequest returns OAuth 2 access and refresh tokens, given the right details:
//
//	basically, a `refresh token` from `AccessTokenRequest`
func RefreshTokenRequest(data utils.H, repositories *repository.RepositoryManager) (utils.H, error) {
	var user models.User
	var client models.Client

	var token string
	var scope string

	if data["refresh_token"] == nil || data["scope"] == nil || data["client"] == nil {
		return shared.InvalidRequestResult("")
	}

	token = data["refresh_token"].(string)
	scope = data["scope"].(string)
	client = data["client"].(models.Client)

	refreshSession := repositories.Sessions().FindByToken(token, models.RefreshToken)
	defer repositories.Sessions().Invalidate(&refreshSession)
	if refreshSession.IsNewRecord() {
		return shared.InvalidGrantResult("")
	}
	user = refreshSession.User
	if refreshSession.Client.ID != client.ID {
		return shared.InvalidGrantResult("")
	}
	if scope != refreshSession.Scopes {
		return shared.InvalidScopeResult("")
	}

	accessToken := models.Session{
		User:      user,
		Client:    client,
		IP:        refreshSession.IP,
		UserAgent: refreshSession.UserAgent,
		Scopes:    scope,
		TokenType: models.AccessToken,
	}
	repositories.Sessions().Create(&accessToken)
	refreshToken := models.Session{
		User:      user,
		Client:    client,
		IP:        refreshSession.IP,
		UserAgent: refreshSession.UserAgent,
		Scopes:    scope,
		TokenType: models.RefreshToken,
	}
	repositories.Sessions().Create(&refreshToken)

	if accessToken.IsNewRecord() || refreshToken.IsNewRecord() {
		return shared.ServerErrorResult("")
	}

	return utils.H{
		"access_token":  accessToken.Token,
		"token_type":    "Bearer",
		"expires_in":    accessToken.ExpiresIn,
		"refresh_token": refreshToken.Token,
	}, nil
}
