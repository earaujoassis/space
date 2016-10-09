package mailer

import (
    "html/template"
    "fmt"
    "bytes"
)

func CreateMessage(templateName string, data interface{}) string {
    templateName = fmt.Sprintf("services/mailer/templates/%s", templateName)
    parser, _ := template.ParseFiles("services/mailer/templates/default.html", templateName)
    buffer := new(bytes.Buffer)
    parser.Execute(buffer, data)
    return buffer.String()
}
