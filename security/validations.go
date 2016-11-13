package security

import (
    "regexp"
)

func ValidUUID(uuid string) bool {
    r := regexp.MustCompile("[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[8|9|aA|bB][a-f0-9]{3}-[a-f0-9]{12}")
    return r.MatchString(uuid)
}

func ValidRandomString(random string) bool {
    r := regexp.MustCompile("[a-zA-Z0-9]+")
    return r.MatchString(random)
}

func ValidToken(token string) bool {
    return ValidRandomString(token)
}

func ValidEmail(email string) bool {
    r := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
    return r.MatchString(email)
}
