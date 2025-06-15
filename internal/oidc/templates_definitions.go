package oidc

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
)

func createCustomRender() multitemplate.Render {
	render := multitemplate.New()
	render.AddFromFiles("formt_post.error", "internal/oidc/templates/form_post.error.html")
	render.AddFromFiles("formt_post.id_token.success", "internal/oidc/templates/form_post.id_token.success.html")
	render.AddFromFiles("formt_post.code.success", "internal/oidc/templates/form_post.code.success.html")
	return render
}
