package tasks

import (
    "fmt"
    "bufio"
    "os"
    "strings"

    "github.com/earaujoassis/space/services"
)

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
    fmt.Print("Client URI redirect: ")
    clientURI, _ := reader.ReadString('\n')
    clientURI = strings.Trim(clientURI, "\n")

    client := services.CreateNewClient(clientName,
        clientDescription,
        clientScope,
        clientURI)
    fmt.Println("A new client application was created")
    fmt.Println("Client key: ", client.Key)
    fmt.Println("Client secret: ", client.Secret)
}

func UpdateClient() {

}
