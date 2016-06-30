package main

import (
    "net/http"
    "fmt"
    "github.com/gin-gonic/gin"
)

func Authentication(router *gin.RouterGroup) {
    users := router.Group("/users")
    {
        users.GET("/", func(c *gin.Context) {
            c.String(http.StatusOK, "not implemented")
        })
        users.POST("/", func(c *gin.Context) {
            dataStore := GetDataStoreConnection()
            /*
             * TODO There should be an actionToken prior to any action.
             * An actionToken should be created whenever a client (web or remote client)
             * attempts a creational action (like creating a user or a session)
             */
            user := User{
                FirstName: c.PostForm("firstName"),
                LastName: c.PostForm("lastName"),
                Username: c.PostForm("username"),
                Email: c.PostForm("email"),
                Passphrase: c.PostForm("password"),
            }
            if dataStore.NewRecord(user.Client) {
                user.Client = FindOrCreateClient("Jupiter")
            }
            if dataStore.NewRecord(user.Language) {
                user.Language = FindOrCreateLanguage("English", "en-US")
            }
            result := dataStore.Create(&user)
            if count := result.RowsAffected; count < 1 {
                c.JSON(406, gin.H{
                    "_status":  "error",
                    "_message": "User was not created",
                    "_self": c.Request.RequestURI,
                    "error": fmt.Sprintf("%v", result.GetErrors()),
                    "user": user,
                })
            } else {
                c.JSON(200, gin.H{
                    "_status":  "created",
                    "_message": "User was created",
                    "_self": c.Request.RequestURI,
                    "user": user,
                })
            }
        })
        users.GET("/:id", func(c *gin.Context) {
            //id := c.Param("id")
            c.String(http.StatusOK, "not implemented")
        })
        users.PUT("/:id", func(c *gin.Context) {
            //id := c.Param("id")
            c.String(http.StatusOK, "not implemented")
        })
        users.DELETE("/:id", func(c *gin.Context) {
            //id := c.Param("id")
            c.String(http.StatusOK, "not implemented")
        })
    }
}
