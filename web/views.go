package web

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/renders/multitemplate"
)

func createCustomRender() multitemplate.Render {
    render := multitemplate.New()
    render.AddFromFiles("satellite", "web/templates/default.html", "web/templates/satellite.html")
    return render
}

func ExposeRoutes(router *gin.Engine) {
    router.LoadHTMLGlob("web/templates/*.html")
    router.HTMLRender = createCustomRender()
    router.Static("/public", "./web/public")
    views := router.Group("/")
    {
        views.GET("/", func(c *gin.Context) {
            c.HTML(http.StatusOK, "satellite", gin.H{
                "Title": "",
                "Satellite": "ganymede",
            })
        })
        views.GET("/signup", func(c *gin.Context) {
            c.HTML(http.StatusOK, "satellite", gin.H{
                "Title": " - Sign up",
                "Satellite": "io",
            })
        })
    }
}
