package logger

import (
    "fmt"
    "time"

    "github.com/earaujoassis/space/services/mailer"
    "github.com/earaujoassis/space/config"
    "github.com/earaujoassis/space/utils"
)

// LogAction is used to Log actions throughout the application;
//      using e-mail messages (production-only) or stdout (development-only)
func LogAction(name string, data utils.H) {
    if config.IsEnvironment("production") {
        switch name {
        case "user.created":
            data["Year"] = time.Now().Year()
            message := mailer.CreateMessage("user.created.html", data)
            mailer.SendEmail("Welcome to quatroLABS services", message, data["Email"].(string))
        case "user.update.password":
            data["Year"] = time.Now().Year()
            message := mailer.CreateMessage("user.update.password.html", data)
            mailer.SendEmail("A magic link to update your password was requested at quatroLABS", message, data["Email"].(string))
        case "user.update.secrets":
            data["Year"] = time.Now().Year()
            message := mailer.CreateMessage("user.update.secrets.html", data)
            mailer.SendEmail("A magic link to recreat your recovery code and secret code generator was requested at quatroLABS", message, data["Email"].(string))
        case "session.created":
            data["Year"] = time.Now().Year()
            message := mailer.CreateMessage("session.created.html", data)
            mailer.SendEmail("A new session created at quatroLABS", message, data["Email"].(string))
        case "session.magic":
            data["Year"] = time.Now().Year()
            message := mailer.CreateMessage("session.magic.html", data)
            mailer.SendEmail("A magic link for a new session was requested at quatroLABS", message, data["Email"].(string))
        }
    }
    if config.IsEnvironment("development") {
        fmt.Printf("[logger] Action `%s` with data `%v`\n", name, data)
    }
}
