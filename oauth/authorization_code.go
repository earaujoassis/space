package oauth

import (
    "strings"
    "encoding/base64"

    "github.com/earaujoassis/space/utils"
    "github.com/earaujoassis/space/datastore"
    "github.com/earaujoassis/space/models"
)

func AuthorizationCodeGrant(data utils.H) (utils.H, error) {
    var redirectURI string
    var scope string
    var state string

    var ip string
    var userAgent string

    var user models.User
    var client models.Client

    if data["redirect_uri"] == nil || data["user"] == nil || data["client"] == nil {
        return invalidRequestResult(state)
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

    redirectURIBytes, _ := base64.StdEncoding.DecodeString(data["redirect_uri"].(string))
    redirectURI = string(redirectURIBytes)
    client = data["client"].(models.Client)
    user = data["user"].(models.User)

    if data["scope"] != nil {
        scope = data["scope"].(string)
    }

    if !strings.Contains(client.RedirectURI, redirectURI) {
        return accessDeniedResult(state)
    }

    if scope != "" && !strings.Contains(client.Scopes, scope) {
        scope = models.PublicScope
    }

    session := models.Session{
        User: user,
        Client: client,
        Ip: ip,
        UserAgent: userAgent,
        Scopes: scope,
        TokenType: models.GrantToken,
    }
    dataStore := datastore.GetDataStoreConnection()
    result := dataStore.Create(&session)
    if count := result.RowsAffected; count > 0 {
        return utils.H{
            "code": session.Token,
            "state": state,
        }, nil
    } else {
        return serverErrorResult(state)
    }
}
