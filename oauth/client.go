package oauth

import (
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/utils"
)

func ClientAuthentication(authorizationHeader string) models.Client {
    key, secret := utils.BasicAuthDecode(authorizationHeader)
    return services.ClientAuthentication(key, secret)
}
