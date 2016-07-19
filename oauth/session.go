package oauth

import (
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
)

func SessionAuthentication(token string) models.Session {
    return services.FindSessionByToken(token, models.AccessToken)
}
