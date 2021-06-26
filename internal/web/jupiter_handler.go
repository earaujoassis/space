package web

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/sessions"

    "github.com/earaujoassis/space/internal/feature"
    "github.com/earaujoassis/space/internal/models"
    "github.com/earaujoassis/space/internal/services"
    "github.com/earaujoassis/space/internal/utils"
)

func jupiterHandler(c *gin.Context) {
    session := sessions.Default(c)
    userPublicID := session.Get("userPublicID")
    if userPublicID == nil {
        c.Redirect(http.StatusFound, "/signin")
        return
    }
    client := services.FindOrCreateClient(services.DefaultClient)
    user := services.FindUserByPublicID(userPublicID.(string))
    actionToken := services.CreateAction(user, client,
        c.Request.RemoteAddr,
        c.Request.UserAgent(),
        models.ReadWriteScope,
        models.NotSpecialAction,
    )
    c.HTML(http.StatusOK, "satellite", utils.H{
        "Title": " - Mission Control",
        "Satellite": "europa",
        "Internal": true,
        "Data": utils.H {
            "action_token": actionToken.Token,
            "user_id": user.UUID,
            "feature.gates": utils.H{
                "user.adminify": feature.IsActive("user.adminify"),
            },
        },
    })
}
