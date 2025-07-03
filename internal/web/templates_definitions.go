package web

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
)

func createCustomRender() multitemplate.Render {
	render := multitemplate.New()
	render.AddFromFiles("satellite", "web/templates/default.html", "web/templates/satellite.html")
	render.AddFromFiles("user.update_secrets", "web/templates/default.html", "web/templates/user.update_secrets.html")
	render.AddFromFiles("error.generic", "web/templates/default.html", "web/templates/error.generic.html")
	render.AddFromFiles("error.email_confirmation", "web/templates/default.html", "web/templates/error.email_confirmation.html")
	render.AddFromFiles("error.password_update", "web/templates/default.html", "web/templates/error.password_update.html")
	render.AddFromFiles("error.secrets_update", "web/templates/default.html", "web/templates/error.secrets_update.html")
	render.AddFromFiles("error.authorization", "web/templates/default.html", "web/templates/error.authorization.html")
	render.AddFromFiles("error.not_found", "web/templates/default.html", "web/templates/error.not_found.html")
	render.AddFromFiles("error.internal", "web/templates/default.html", "web/templates/error.internal.html")
	return render
}
