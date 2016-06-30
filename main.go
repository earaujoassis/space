package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/gin-gonic/contrib/renders/multitemplate"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("templates/*.html")
    router.HTMLRender = createCustomRender()
    router.Static("/public", "./public")
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
        views.GET("/success", func(c *gin.Context) {
            c.HTML(http.StatusOK, "success", gin.H{
                "Title": " - Success!",
                "Satellite": "io",
            })
        })
    }
    restApi := router.Group("/api/v1")
    Authentication(restApi)
    GetDataStoreConnection().AutoMigrate(&Client{}, &Language{}, &User{}, &Session{})
    router.Run(fmt.Sprintf(":%v", GetConfig("http.port")))
}

func createCustomRender() multitemplate.Render {
    render := multitemplate.New()
    render.AddFromFiles("satellite", "templates/default.html", "templates/satellite.html")
    render.AddFromFiles("success", "templates/default.html", "templates/success.html")
    return render
}
