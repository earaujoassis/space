package tasks

import (
    "fmt"
    "bufio"
    "os"
    "strings"

    "github.com/earaujoassis/space/internal/models"
    "github.com/earaujoassis/space/internal/services"
)

// CreateClient task is used to create a new client application
func CreateClient() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Client name: ")
    clientName, _ := reader.ReadString('\n')
    clientName = strings.Trim(clientName, "\n")
    fmt.Print("Client description: ")
    clientDescription, _ := reader.ReadString('\n')
    clientDescription = strings.Trim(clientDescription, "\n")
    fmt.Print("Client scope: ")
    clientScope, _ := reader.ReadString('\n')
    clientScope = strings.Trim(clientScope, "\n")
    fmt.Print("Client canonical URI: ")
    canonicalURI, _ := reader.ReadString('\n')
    canonicalURI = strings.Trim(canonicalURI, "\n")
    fmt.Print("Client URI redirect: ")
    redirectURI, _ := reader.ReadString('\n')
    redirectURI = strings.Trim(redirectURI, "\n")

    clientSecret := models.GenerateRandomString(64)
    client := services.CreateNewClient(clientName,
        clientDescription,
        clientSecret,
        clientScope,
        canonicalURI,
        redirectURI)
    if client.ID == 0 {
        fmt.Println("There's a error and the client was not created")
    } else {
        fmt.Println("A new client application was created")
        fmt.Println("Client key:", client.Key)
        fmt.Println("Client secret:", clientSecret)
    }
}
