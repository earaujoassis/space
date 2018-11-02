package oauth

import (
    "github.com/earaujoassis/space/models"
    "github.com/earaujoassis/space/services"
    "github.com/earaujoassis/space/utils"
)

// ClientAuthentication authenticates a client application, extracting the key-secret pair;
//      and returns a client entry/model, given the key-secret pair
func ClientAuthentication(authorizationHeader string) models.Client {
    key, secret := utils.BasicAuthDecode(authorizationHeader)
    return services.ClientAuthentication(key, secret)
}
