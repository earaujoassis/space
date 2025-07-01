package oauth

import (
	"golang.org/x/exp/slices"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/repository"
	"github.com/earaujoassis/space/internal/shared"
)

type AccessTokenParams struct {
	GrantType   string
	Code        string
	RedirectURI string
	Client      models.Client
}

type AccessTokenResult struct {
	AccessToken  string
	TokenType    string
	ExpiresIn    int64
	RefreshToken string
}

// AccessTokenRequest returns OAuth 2 access and refresh tokens, given the right details:
//
//	basically, a `code` from `AuthorizationCodeGrant`
func AccessTokenRequest(params AccessTokenParams, repositories *repository.RepositoryManager) (*AccessTokenResult, *shared.RequestError) {
	var client models.Client

	var code string
	var redirectURI string

	if params.Code == "" || params.RedirectURI == "" || params.Client.IsNewRecord() {
		return nil, shared.InvalidRequestResult("")
	}

	redirectURI = params.RedirectURI
	code = params.Code
	client = params.Client

	authorizationSession := repositories.Sessions().FindByToken(code, models.GrantToken)
	defer repositories.Sessions().Invalidate(&authorizationSession)
	if authorizationSession.IsNewRecord() {
		return nil, shared.InvalidGrantResult("")
	}
	user := authorizationSession.User
	if authorizationSession.Client.ID != client.ID {
		return nil, shared.InvalidGrantResult("")
	}
	if !slices.Contains(authorizationSession.Client.RedirectURI, redirectURI) {
		return nil, shared.InvalidGrantResult("")
	}

	accessToken := models.Session{
		User:      user,
		Client:    client,
		IP:        authorizationSession.IP,
		UserAgent: authorizationSession.UserAgent,
		Scopes:    authorizationSession.Scopes,
		TokenType: models.AccessToken,
	}
	repositories.Sessions().Create(&accessToken)
	refreshToken := models.Session{
		User:      user,
		Client:    client,
		IP:        authorizationSession.IP,
		UserAgent: authorizationSession.UserAgent,
		Scopes:    authorizationSession.Scopes,
		TokenType: models.RefreshToken,
	}
	repositories.Sessions().Create(&refreshToken)

	if accessToken.IsNewRecord() || refreshToken.IsNewRecord() {
		return nil, shared.ServerErrorResult("")
	}

	return &AccessTokenResult{
		AccessToken:  accessToken.Token,
		TokenType:    "Bearer",
		ExpiresIn:    accessToken.ExpiresIn,
		RefreshToken: refreshToken.Token,
	}, nil
}
