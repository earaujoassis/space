package helpers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type OAuthTestClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewOAuthTestClient(baseURL string) *OAuthTestClient {
	jar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Error creating cookiejar: %s", err)
	}
	return &OAuthTestClient{
		baseURL:    baseURL,
		httpClient: &http.Client{
			Jar: jar,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func (c *OAuthTestClient) LoginUser(holder, password, passcode string) *TestResponse {
	requestUrl := c.baseURL + "/api/sessions/create"
	formData := url.Values{}
	formData.Set("holder", holder)
	formData.Set("password", password)
	formData.Set("passcode", passcode)
	encoded := formData.Encode()
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(encoded))
	request.Header.Set("X-Requested-By", "SpaceApi")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := c.httpClient.Do(request)
	return parseResponse(response, err)
}

func (c *OAuthTestClient) GetAuthorization(responseType, clientID, redirectURI string) *TestResponse {
	params := url.Values{}
	params.Set("response_type", responseType)
	params.Set("client_id", clientID)
	params.Set("redirect_uri", redirectURI)
	params.Set("state", "test-state")

	authURL := fmt.Sprintf("%s/oauth/authorize?%s", c.baseURL, params.Encode())
	resp, err := c.httpClient.Get(authURL)
	return parseResponse(resp, err)
}
