package main

import (
    "log"
    "net/http"
    "os"
    "fmt"

    "./energy"
)

func main() {
    router := energy.NewApplication(routes)
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    portString := fmt.Sprintf(":%s", port)

    log.Printf("Starting server at port %s", port)
    log.Fatal(http.ListenAndServe(portString, router))
}
