package utils

import (
    "encoding/base64"
    "strings"
)

func BasicAuthEncode(username, password string) string {
    token := username + ":" + password
    return base64.StdEncoding.EncodeToString([]byte(token))
}

func BasicAuthDecode(token string) (string, string) {
    bytes, _ := base64.StdEncoding.DecodeString(token)
    values := strings.Split(string(bytes), ":")
    return values[0], values[1]
}
