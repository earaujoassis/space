package oauth

import (
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
)

// AccessAuthentication obtains a Session entry (typed as an `Access Token`) through
//      its token string
func AccessAuthentication(token string) models.Session {
    return services.FindSessionByToken(token, models.AccessToken)
}
