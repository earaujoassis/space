package oidc

import (
	"golang.org/x/exp/slices"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/repository"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

type AuthorizationCodeParams struct {
	ResponseType string
	Client       models.Client
	User         models.User
	IP           string
	UserAgent    string
	RedirectURI  string
	Scope        string
	State        string
	Nonce        string
	ResponseMode string
}

type AuthorizationCodeResult struct {
	Code  string
	Scope string
	State string
}

// AuthorizationCodeGrant returns an OIDC authorization code grant, given the right details
func AuthorizationCodeGrant(params AuthorizationCodeParams, repositories *repository.RepositoryManager) (*AuthorizationCodeResult, *shared.RequestError) {
	var redirectURI string
	var scope string
	var state string

	var ip string
	var userAgent string

	var user models.User
	var client models.Client

	if params.RedirectURI == "" || params.User.IsNewRecord() || params.Client.IsNewRecord() {
		return nil, shared.InvalidRequestResult(state)
	}

	state = params.State
	ip = params.IP
	userAgent = params.UserAgent
	redirectURI = params.RedirectURI
	client = params.Client
	user = params.User
	scope = params.Scope

	if !slices.Contains(client.RedirectURI, redirectURI) {
		return nil, shared.InvalidRequestResult(state)
	}

	if scope != "" && !client.HasRequestedScopes(utils.Scopes(scope)) && !client.HasScope(models.OpenIDScope) {
		return nil, shared.InvalidScopeResult(state)
	}

	session := models.Session{
		User:      user,
		Client:    client,
		IP:        ip,
		UserAgent: userAgent,
		Scopes:    scope,
		TokenType: models.GrantToken,
	}
	repositories.Sessions().Create(&session)
	nonce := models.Nonce{
		ClientKey: client.Key,
		Code:      session.Token,
		Nonce:     params.Nonce,
	}
	if nonce.Nonce != "" {
		if !nonce.IsValid() {
			return nil, shared.InvalidRequestResult(state)
		}
		if ok := repositories.Nonces().Create(nonce); !ok {
			return nil, shared.InvalidRequestResult(state)
		}
	}
	if session.IsSavedRecord() {
		return &AuthorizationCodeResult{
			Code:  session.Token,
			State: state,
			Scope: scope,
		}, nil
	}

	return nil, shared.ServerErrorResult(state)
}
