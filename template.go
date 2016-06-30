package main

import (
    "github.com/gin-gonic/contrib/renders/multitemplate"
)

func CreateCustomRender() multitemplate.Render {
    r := multitemplate.New()
    r.AddFromFiles("ganymede", "templates/default.html", "templates/ganymede.html")

    return r
}
