package helpers

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

type TestResponse struct {
	StatusCode int
	Body       string
	Headers    http.Header
	Location   string
	JSON       map[string]interface{}
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

	return result
}
