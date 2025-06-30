package oidc

import (
	"slices"
	"strings"

	"github.com/earaujoassis/space/internal/models"
	"github.com/earaujoassis/space/internal/repository"
	"github.com/earaujoassis/space/internal/shared"
	"github.com/earaujoassis/space/internal/utils"
)

type ImplicitFlowIDTokenParams struct {
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
	Issuer       string
}

type ImplicitFlowIDTokenResult struct {
	IDToken string
	State   string
}

// ImplicitFlowIDToken returns an OIDC id_token grant, given the right details
func ImplicitFlowIDToken(params ImplicitFlowIDTokenParams, repositories *repository.RepositoryManager) (*ImplicitFlowIDTokenResult, *shared.RequestError) {
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
	issuer := params.Issuer

	nonce := models.Nonce{
		ClientKey: client.Key,
		Code:      "",
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

	scope = params.Scope
	if !slices.Contains(client.RedirectURI, redirectURI) {
		return nil, shared.InvalidRequestResult(state)
	}

	if scope != "" && !client.HasRequestedScopes(utils.Scopes(scope)) && !strings.Contains(scope, models.OpenIDScope) {
		return nil, shared.InvalidScopeResult(state)
	}

	idToken := createIDToken(issuer, user.PublicID, client.Key, nonce.Nonce)
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
	if idToken != "" && session.IsSavedRecord() {
		return &ImplicitFlowIDTokenResult{
			IDToken: idToken,
			State:   state,
		}, nil
	}

	return nil, shared.ServerErrorResult(state)
}
