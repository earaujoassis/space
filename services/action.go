package services

import (
    "github.com/earaujoassis/space/models"
)

func CreateAction(user models.User, client models.Client, ip, userAgent, scopes string) models.Action {
    var action models.Action = models.Action{
        User: user,
        Client: client,
        Ip: ip,
        UserAgent: userAgent,
        Scopes: scopes,
    }
    if err := action.Save(); err != nil {
        return models.Action{}
    }
    return action
}

func ActionAuthentication(token string) models.Action {
    var action models.Action = models.RetrieveActionByToken(token)
    if action.UUID != "" && !action.WithinExpirationWindow() {
        action.Delete()
        return models.Action{}
    }
    return action
}

func ActionGrantsReadAbility(action models.Action) bool {
    return action.Scopes == models.ReadScope || action.Scopes == models.ReadWriteScope
}

func ActionGrantsWriteAbility(action models.Action) bool {
    return action.Scopes == models.ReadWriteScope
}
