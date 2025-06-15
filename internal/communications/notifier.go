package communications

import (
	"time"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/mailer"
	"github.com/earaujoassis/space/internal/utils"
)

type Notifier struct {
	cfg *config.Config
}

func NewNotifier(cfg *config.Config) *Notifier {
	return &Notifier{
		cfg: cfg,
	}
}

// Announce is used to communicate actions throughout the application,
//
//	using e-mail messages (production-only) or stdout (development-only)
func (n *Notifier) Announce(name string, data utils.H) {
	if n.cfg.Environment == "production" {
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
	if n.cfg.Environment == "development" {
		logs.Propagatef(logs.Info, "Action `%s` with data `%v`\n", name, data)
	}
}
