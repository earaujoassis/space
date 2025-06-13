package integration

import (
	"fmt"
)

const (
	Code              string = "code"
	AuthorizationCode string = "authorization_code"
	Token             string = "token"
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
