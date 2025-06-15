package integration

import (
	"fmt"
	"strings"
)

const (
	Code              string = "code"
	AuthorizationCode string = "authorization_code"
	Token             string = "token"
	IdToken           string = "id_token"
)

func interfaceSliceToStringSlice(slice []interface{}) []string {
	result := make([]string, len(slice))
	for i, v := range slice {
		if str, ok := v.(string); ok {
			result[i] = str
		} else {
			result[i] = fmt.Sprintf("%v", v)
		}
	}
	return result
}

func jsonValueAsSingleString(value interface{}) string {
	return strings.Join(interfaceSliceToStringSlice(value.([]interface{})), " ")
}
