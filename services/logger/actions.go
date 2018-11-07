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
            mailer.SendEmail("Welcome to QuatroLabs services", message, data["Email"].(string))
        case "session.created":
            data["Year"] = time.Now().Year()
            message := mailer.CreateMessage("session.created.html", data)
            mailer.SendEmail("A new session created at QuatroLabs", message, data["Email"].(string))
        }
    }
    if config.IsEnvironment("development") {
        fmt.Printf("[logger] Action `%s` with data `%v`\n", name, data)
    }
}
