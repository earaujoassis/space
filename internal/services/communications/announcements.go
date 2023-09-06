package communications

import (
	"time"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/services/mailer"
	"github.com/earaujoassis/space/internal/utils"
)

// Announce is used to communicate actions throughout the application,
//
//	using e-mail messages (production-only) or stdout (development-only)
func Announce(name string, data utils.H) {
	if config.IsEnvironment("production") {
		var err error
		switch name {
		case "user.created":
			data["Year"] = time.Now().Year()
			message := mailer.CreateMessage("user.created.html", data)
			err = mailer.SendEmail(
				"Welcome to quatroLABS services",
				message,
				data["Email"].(string))
		case "user.update.password":
			data["Year"] = time.Now().Year()
			message := mailer.CreateMessage("user.update.password.html", data)
			err = mailer.SendEmail(
				"A magic link to update your password was requested at quatroLABS",
				message,
				data["Email"].(string))
		case "user.update.secrets":
			data["Year"] = time.Now().Year()
			message := mailer.CreateMessage("user.update.secrets.html", data)
			err = mailer.SendEmail(
				"A magic link to recreat your recovery code and secret code generator was requested at quatroLABS",
				message,
				data["Email"].(string))
		case "session.created":
			data["Year"] = time.Now().Year()
			message := mailer.CreateMessage("session.created.html", data)
			err = mailer.SendEmail(
				"A new session created at quatroLABS",
				message,
				data["Email"].(string))
		case "session.magic":
			data["Year"] = time.Now().Year()
			message := mailer.CreateMessage("session.magic.html", data)
			err = mailer.SendEmail(
				"A magic link for a new session was requested at quatroLABS",
				message,
				data["Email"].(string))
		}
		if err != nil {
			logs.Propagatef(
				logs.Critical,
				"Critical error sending announcement: action `%s` with data `%v`\n",
				name,
				data)
		}

	}
	if config.IsEnvironment("development") {
		logs.Propagatef(logs.Info, "Action `%s` with data `%v`\n", name, data)
	}
}
