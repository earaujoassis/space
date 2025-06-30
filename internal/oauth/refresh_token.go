package oauth

import (
	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/repository"
	"github.com/earaujoassis/space/internal/shared"
)

type RefreshTokenParams struct {
	GrantType    string
	RefreshToken string
	Scope        string
	Client       models.Client
}

type RefreshTokenResult struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int64
	RefreshToken string
}

// RefreshTokenRequest returns OAuth 2 access and refresh tokens, given the right details:
//
//	basically, a `refresh token` from `AccessTokenRequest`
func RefreshTokenRequest(params RefreshTokenParams, repositories *repository.RepositoryManager) (*RefreshTokenResult, *shared.RequestError) {
	var user models.User
	var client models.Client

	var token string
	var scope string

	if params.RefreshToken == "" || params.Scope == "" || params.Client.IsNewRecord() {
		return nil, shared.InvalidRequestResult("")
	}

	token = params.RefreshToken
	scope = params.Scope
	client = params.Client

	refreshSession := repositories.Sessions().FindByToken(token, models.RefreshToken)
	defer repositories.Sessions().Invalidate(&refreshSession)
	if refreshSession.IsNewRecord() {
		return nil, shared.InvalidGrantResult("")
	}
	user = refreshSession.User
	if refreshSession.Client.ID != client.ID {
		return nil, shared.InvalidGrantResult("")
	}
	if scope != refreshSession.Scopes {
		return nil, shared.InvalidScopeResult("")
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
		return nil, shared.ServerErrorResult("")
	}

	return &RefreshTokenResult{
		AccessToken:  accessToken.Token,
		TokenType:    "Bearer",
		ExpiresIn:    accessToken.ExpiresIn,
		RefreshToken: refreshToken.Token,
	}, nil
}
