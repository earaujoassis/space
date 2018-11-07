package mailer

import (
    "html/template"
    "fmt"
    "bytes"
)

// CreateMessage creates an e-mail message to be sent through the mailer service
func CreateMessage(templateName string, data interface{}) string {
    templateName = fmt.Sprintf("services/mailer/templates/%s", templateName)
    parser, _ := template.ParseFiles("services/mailer/templates/default.html", templateName)
    buffer := new(bytes.Buffer)
    parser.Execute(buffer, data)
    return buffer.String()
}
