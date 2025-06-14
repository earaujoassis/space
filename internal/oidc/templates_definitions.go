package oidc

import (
	"github.com/gin-gonic/contrib/renders/multitemplate"
)

func createCustomRender() multitemplate.Render {
	render := multitemplate.New()
	render.AddFromFiles("formt_post", "internal/oidc/templates/form_post.html")
	return render
}
