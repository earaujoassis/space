package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type TestResponse struct {
	StatusCode int
	Body       string
	Headers    http.Header
	Location   string
	JSON       map[string]interface{}
	Query      map[string]string
}

func parseQueryString(rawURL string) map[string]string {
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

func parseResponse(response *http.Response, err error) *TestResponse {
	if err != nil {
		return &TestResponse{StatusCode: 0, Body: err.Error()}
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	bodyStr := string(bodyBytes)

	result := &TestResponse{
		StatusCode: response.StatusCode,
		Body:       bodyStr,
		Headers:    response.Header,
		Location:   response.Header.Get("Location"),
	}

	if strings.Contains(response.Header.Get("Content-Type"), "application/json") {
		var jsonData map[string]interface{}
		if json.Unmarshal(bodyBytes, &jsonData) == nil {
			result.JSON = jsonData
		}
	}

	if result.Location != "" {
		result.Query = parseQueryString(result.Location)
	}

	return result
}

func (r *TestResponse) HasKeyInQuery(key string) bool {
	_, ok := r.Query[key]
	return ok
}
