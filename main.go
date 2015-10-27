package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    . "github.com/earaujoassis/space/energy"
    "github.com/jinzhu/gorm"
)

var Database gorm.DB

func main() {
    var err error
    router := NewApplication(routes)
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    portString := fmt.Sprintf(":%s", port)

    databaseUrl := fmt.Sprintf(
        "host=%s port=5432 user=%s dbname=%s sslmode=disable",
        "192.168.44.88",
        "postgres",
        "space_development",
    )
    Database, err = gorm.Open("postgres", databaseUrl)
    if err != nil {
        log.Fatal("Could not connect to database")
    }
    defer Database.Close()

    log.Printf("Starting server at port %s", port)
    log.Fatal(http.ListenAndServe(portString, router))
}
