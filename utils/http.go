package utils

import (
    "encoding/base64"
    "strings"
)

// BasicAuthEncode encodes a key-secret pair to be used in a HTTP Basic Authentication
//      strategy (HTTP Request Header)
func BasicAuthEncode(key, secret string) string {
    token := key + ":" + secret
    return base64.StdEncoding.EncodeToString([]byte(token))
}

// BasicAuthDecode decodes a key-secret pair from a token string, originally used in a
//      HTTP Basic Authentication strategy (HTTP Request Header)
func BasicAuthDecode(token string) (string, string) {
    bytes, _ := base64.StdEncoding.DecodeString(token)
    values := strings.Split(string(bytes), ":")
    return values[0], values[1]
}
