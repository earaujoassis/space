package utils

import (
	"net/url"
)

func ParseQueryString(rawURL string) map[string]string {
    result := make(map[string]string)

    u, err := url.Parse(rawURL)
    if err != nil {
        return result
    }

    for key, values := range u.Query() {
        if len(values) > 0 {
            result[key] = values[0]
        }
    }

    return result
}

func ParseFragment(rawURL string) map[string]string {
	result := make(map[string]string)

	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return result
	}

	fragment := parsedURL.Fragment
	if fragment == "" {
		return result
	}

	values, err := url.ParseQuery(fragment)
	if err != nil {
		return result
	}

	for key, vals := range values {
		if len(vals) > 0 {
			result[key] = vals[0]
		}
	}

	return result
}
