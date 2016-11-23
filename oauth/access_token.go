package oauth

import (
    "strings"

    "github.com/earaujoassis/space/utils"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/models"
)

func AccessTokenRequest(data utils.H) (utils.H, error) {
    var user models.User
    var client models.Client

    var code string
    var redirectURI string

    if data["code"] == nil || data["redirect_uri"] == nil || data["client"] == nil {
        return invalidRequestResult("")
    }

    redirectURI = data["redirect_uri"].(string)
    code = data["code"].(string)
    client = data["client"].(models.Client)

    authorizationSession := services.FindSessionByToken(code, models.GrantToken)
    defer services.InvalidateSession(authorizationSession)
    if authorizationSession.ID == 0 {
        return invalidGrantResult("")
    }
    user = authorizationSession.User
    user = services.FindUserByPublicId(user.PublicId)
    if authorizationSession.Client.ID != client.ID {
        return invalidGrantResult("")
    }
    if !strings.Contains(authorizationSession.Client.RedirectURI, redirectURI) {
        return invalidGrantResult("")
    }

    accessToken := services.CreateSession(user,
        client,
        authorizationSession.Ip,
        authorizationSession.UserAgent,
        authorizationSession.Scopes,
        models.AccessToken)
    refreshToken := services.CreateSession(user,
        client,
        authorizationSession.Ip,
        authorizationSession.UserAgent,
        authorizationSession.Scopes,
        models.RefreshToken)

    if accessToken.ID == 0 || refreshToken.ID == 0 {
        return serverErrorResult("")
    }

    return utils.H{
        "user_id": user.PublicId,
        "access_token": accessToken.Token,
        "token_type": "Bearer",
        "expires_in": accessToken.ExpiresIn,
        "refresh_token": refreshToken.Token,
        "scope": authorizationSession.Scopes,
    }, nil
}

func RefreshTokenRequest(data utils.H) (utils.H, error) {
    var user models.User
    var client models.Client

    var token string
    var scope string

    if data["refresh_token"] == nil || data["scope"] == nil || data["client"] == nil {
        return invalidRequestResult("")
    }

    token = data["refresh_token"].(string)
    scope = data["scope"].(string)
    client = data["client"].(models.Client)

    refreshSession := services.FindSessionByToken(token, models.RefreshToken)
    defer services.InvalidateSession(refreshSession)
    if refreshSession.ID == 0 {
        return invalidGrantResult("")
    }
    user = refreshSession.User
    user = services.FindUserByPublicId(user.PublicId)
    if refreshSession.Client.ID != client.ID {
        return invalidGrantResult("")
    }
    if scope != refreshSession.Scopes {
        return invalidScopeResult("")
    }

    accessToken := services.CreateSession(user,
        client,
        refreshSession.Ip,
        refreshSession.UserAgent,
        scope,
        models.AccessToken)
    refreshToken := services.CreateSession(user,
        client,
        refreshSession.Ip,
        refreshSession.UserAgent,
        scope,
        models.RefreshToken)

    if accessToken.ID == 0 || refreshToken.ID == 0 {
        return serverErrorResult("")
    }

    return utils.H{
        "user_id": user.PublicId,
        "access_token": accessToken.Token,
        "token_type": "Bearer",
        "expires_in": accessToken.ExpiresIn,
        "refresh_token": refreshToken.Token,
        "scope": refreshSession.Scopes,
    }, nil
}
