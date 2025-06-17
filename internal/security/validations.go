package security

import (
	"regexp"
)

// ValidUUID checks if `uuid` is a valid UUID-v4 string
func ValidUUID(uuid string) bool {
	validChars := regexp.MustCompile("^[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|aA|bB][a-f0-9]{3}-[a-f0-9]{12}$")
	return validChars.MatchString(uuid)
}

// ValidBase64 checks if `encoded` is a valid base64 string
func ValidBase64(encoded string) bool {
	validChars := regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=|[A-Za-z0-9+/]{4})$")
	return validChars.MatchString(encoded)
}

// ValidRandomString checks if `random` is a valid random string
func ValidRandomString(random string) bool {
	validChars := regexp.MustCompile("^[a-zA-Z0-9]+$")
	return validChars.MatchString(random)
}

// ValidToken checks if `token` is a valid token/random string
func ValidToken(token string) bool {
	return ValidRandomString(token)
}

// ValidNonce checks if `nonce` is a valid nonce/random string
func ValidNonce(nonce string) bool {
	if len(nonce) < 8 || len(nonce) > 128 {
		return false
	}

	validChars := regexp.MustCompile(`^[a-zA-Z0-9\-_\.~]+$`)
	return validChars.MatchString(nonce)
}

// ValidEmail checks if `email` is a valid e-mail address
func ValidEmail(email string) bool {
	validChars := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return validChars.MatchString(email)
}
