package mailer

import (
	"time"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/logs"
	"github.com/earaujoassis/space/internal/utils"
)

// Announce is used to communicate actions throughout the application,
//
//	using e-mail messages (production-only) or stdout (development-only)
func (m *Mailer) AnnounceMessage(name string, data utils.H) error {
	switch m.cfg.Environment {
	case config.Production:
		var err error
		switch name {
		case "client.token_introspection":
			data["Year"] = time.Now().Year()
			message := CreateMessage("client.token_introspection.html", data)
			err = m.SendEmail(
				"Client application performed token introspection at quatroLABS",
				message,
				data["Email"].(string))
		case "client.userinfo_introspection":
			data["Year"] = time.Now().Year()
			message := CreateMessage("client.userinfo_introspection.html", data)
			err = m.SendEmail(
				"Client application performed userinfo introspection at quatroLABS",
				message,
				data["Email"].(string))
		case "user.authorization_granted":
			data["Year"] = time.Now().Year()
			message := CreateMessage("user.authorization_granted.html", data)
			err = m.SendEmail(
				"A new authorization was granted at quatroLABS",
				message,
				data["Email"].(string))
		case "user.created":
			data["Year"] = time.Now().Year()
			message := CreateMessage("user.created.html", data)
			err = m.SendEmail(
				"Welcome to quatroLABS services",
				message,
				data["Email"].(string))
		case "user.update_password":
			data["Year"] = time.Now().Year()
			message := CreateMessage("user.update_password.html", data)
			err = m.SendEmail(
				"A magic link to update your password was requested at quatroLABS",
				message,
				data["Email"].(string))
		case "user.update_secrets":
			data["Year"] = time.Now().Year()
			message := CreateMessage("user.update_secrets.html", data)
			err = m.SendEmail(
				"A magic link to recreat your recovery code and secret code generator was requested at quatroLABS",
				message,
				data["Email"].(string))
		case "user.email_verification":
			data["Year"] = time.Now().Year()
			message := CreateMessage("user.email_verification.html", data)
			err = m.SendEmail(
				"Please confirm you e-mail address at quatroLABS",
				message,
				data["Email"].(string))
		case "session.created":
			data["Year"] = time.Now().Year()
			message := CreateMessage("session.created.html", data)
			err = m.SendEmail(
				"A new session created at quatroLABS",
				message,
				data["Email"].(string))
		case "session.magic":
			data["Year"] = time.Now().Year()
			message := CreateMessage("session.magic.html", data)
			err = m.SendEmail(
				"A magic link for a new session was requested at quatroLABS",
				message,
				data["Email"].(string))
		}
		if err != nil {
			logs.Propagatef(
				logs.LevelCritical,
				"Critical error sending announcement: action `%s` with data `%v`\n",
				name,
				data)
			return err
		}
	case config.Development:
		logs.Propagatef(logs.LevelInfo, "Action `%s` with data `%v`\n", name, data)
	}

	return nil
}
