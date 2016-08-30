package oauth

import (
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
)

func AccessAuthentication(token string) models.Session {
    return services.FindSessionByToken(token, models.AccessToken)
}

func ActionAuthentication(token string) models.Session {
    return services.FindSessionByToken(token, models.ActionToken)
}
