package logger

import (
    "time"

    "github.com/earaujoassis/space/services/mailer"
    "github.com/earaujoassis/space/utils"
)

func LogAction(name string, data utils.H) {
    switch name {
    case "user.created":
        data["Year"] = time.Now().Year()
        message := mailer.CreateMessage("user.created.html", data)
        mailer.SendEmail("Welcome to QuatroLabs services", message, data["Email"].(string))
    case "session.created":
        data["Year"] = time.Now().Year()
        message := mailer.CreateMessage("session.created.html", data)
        mailer.SendEmail("A new session created at QuatroLabs", message, data["Email"].(string))
    }
}
