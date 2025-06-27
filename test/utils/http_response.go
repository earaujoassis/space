package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/earaujoassis/space/internal/utils"
)

type TestResponse struct {
	StatusCode int
	Body       string
	Headers    http.Header
	Location   string
	JSON       utils.H
	Query      map[string]string
	Fragment   map[string]string
}

func ParseResponse(response *http.Response, err error) *TestResponse {
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
		var jsonData utils.H
		if json.Unmarshal(bodyBytes, &jsonData) == nil {
			result.JSON = jsonData
		}
	}

	if result.Location != "" {
		result.Query = utils.ParseQueryString(result.Location)
		result.Fragment = utils.ParseFragmentString(result.Location)
	}

	return result
}

func (r *TestResponse) HasKeyInQuery(key string) bool {
	_, ok := r.Query[key]
	return ok
}

func (r *TestResponse) HasKeyInFragment(key string) bool {
	_, ok := r.Fragment[key]
	return ok
}

func (r *TestResponse) HasKeyInJSON(key string) bool {
	_, ok := r.JSON[key]
	return ok
}
