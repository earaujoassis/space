package helpers

import (
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"golang.org/x/net/publicsuffix"

	"github.com/earaujoassis/space/test/factory"
	"github.com/earaujoassis/space/test/utils"
)

type OAuthTestClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewOAuthTestClient(baseURL string) *OAuthTestClient {
	options := &cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(options)
	if err != nil {
		log.Fatalf("Error creating cookiejar: %s", err)
	}
	return &OAuthTestClient{
		baseURL: baseURL,
		httpClient: &http.Client{
			Jar: jar,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

func (c *OAuthTestClient) ClearSession() {
	options := &cookiejar.Options{
		PublicSuffixList: publicsuffix.List,
	}
	jar, err := cookiejar.New(options)
	if err != nil {
		log.Fatalf("Error creating cookiejar: %s", err)
	}
	c.httpClient.Jar = jar
}

func (c *OAuthTestClient) StartSession(user *factory.User) {
	c.ClearSession()
	code := user.GenerateCode()
	response := c.LoginUser(user.Model.Username, user.Passphrase, code)
	json := response.JSON
	location := c.baseURL +
		fmt.Sprintf("%s", json["redirect_uri"]) +
		fmt.Sprintf("?client_id=%s", json["client_id"]) +
		fmt.Sprintf("&code=%s", json["code"]) +
		fmt.Sprintf("&grant_type=%s", json["grant_type"]) +
		fmt.Sprintf("&scope=%s", json["scope"]) +
		fmt.Sprintf("&state=%s", json["state"])
	resp, _ := c.httpClient.Get(location)
	u, _ := url.Parse(c.baseURL)
	c.httpClient.Jar.SetCookies(u, resp.Cookies())
}

func (c *OAuthTestClient) HasSessionCookie() bool {
	u, _ := url.Parse(c.baseURL)
	cookies := c.httpClient.Jar.Cookies(u)

	for _, cookie := range cookies {
		if cookie.Name == "space.session" {
			return cookie.Value != ""
		}
	}

	return false
}

func (c *OAuthTestClient) LoginUser(holder, password, passcode string) *utils.TestResponse {
	requestUrl := c.baseURL + "/api/sessions"
	formData := url.Values{}
	formData.Set("holder", holder)
	formData.Set("password", password)
	formData.Set("passcode", passcode)
	encoded := formData.Encode()
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(encoded))
	request.Header.Set("X-Requested-By", "SpaceApi")
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	response, err := c.httpClient.Do(request)
	return utils.ParseResponse(response, err)
}

func (c *OAuthTestClient) GetAuthorize(responseType, clientID, redirectURI, state string) *utils.TestResponse {
	params := url.Values{}
	params.Set("response_type", responseType)
	params.Set("client_id", clientID)
	params.Set("redirect_uri", redirectURI)
	params.Set("scope", "public")
	params.Set("state", state)
	authURL := fmt.Sprintf("%s/oauth/authorize?%s", c.baseURL, params.Encode())
	resp, err := c.httpClient.Get(authURL)
	return utils.ParseResponse(resp, err)
}

func (c *OAuthTestClient) PostAuthorize(responseType, clientID, redirectURI, state string, authorize bool) *utils.TestResponse {
	params := url.Values{}
	params.Set("response_type", responseType)
	params.Set("client_id", clientID)
	params.Set("redirect_uri", redirectURI)
	params.Set("scope", "public")
	params.Set("state", state)
	authURL := fmt.Sprintf("%s/oauth/authorize?%s", c.baseURL, params.Encode())
	if authorize {
		resp, err := c.httpClient.Post(authURL, "application/x-www-form-urlencoded", strings.NewReader(""))
		return utils.ParseResponse(resp, err)
	} else {
		formData := url.Values{}
		formData.Set("access_denied", "true")
		encoded := formData.Encode()
		resp, err := c.httpClient.Post(authURL, "application/x-www-form-urlencoded", strings.NewReader(encoded))
		return utils.ParseResponse(resp, err)
	}
}

func (c *OAuthTestClient) PostToken(clientBasicAuth, grantType string) *utils.TestResponse {
	formData := url.Values{}
	formData.Set("grant_type", grantType)
	encoded := formData.Encode()
	requestUrl := fmt.Sprintf("%s/oauth/token", c.baseURL)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(encoded))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", clientBasicAuth))
	response, err := c.httpClient.Do(request)
	return utils.ParseResponse(response, err)
}

func (c *OAuthTestClient) PostTokenComplete(clientBasicAuth, grantType, code, redirectURI string) *utils.TestResponse {
	formData := url.Values{}
	formData.Set("grant_type", grantType)
	formData.Set("code", code)
	formData.Set("redirect_uri", redirectURI)
	encoded := formData.Encode()
	requestUrl := fmt.Sprintf("%s/oauth/token", c.baseURL)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(encoded))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", clientBasicAuth))
	response, err := c.httpClient.Do(request)
	return utils.ParseResponse(response, err)
}

func (c *OAuthTestClient) PostTokenRefresh(clientBasicAuth, refreshToken, scope string) *utils.TestResponse {
	formData := url.Values{}
	formData.Set("grant_type", "refresh_token")
	formData.Set("refresh_token", refreshToken)
	formData.Set("scope", scope)
	encoded := formData.Encode()
	requestUrl := fmt.Sprintf("%s/oauth/token", c.baseURL)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(encoded))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", clientBasicAuth))
	response, err := c.httpClient.Do(request)
	return utils.ParseResponse(response, err)
}

func (c *OAuthTestClient) PostIntrospect(clientBasicAuth, token string) *utils.TestResponse {
	formData := url.Values{}
	formData.Set("token", token)
	encoded := formData.Encode()
	requestUrl := fmt.Sprintf("%s/oauth/introspect", c.baseURL)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(encoded))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", clientBasicAuth))
	response, err := c.httpClient.Do(request)
	return utils.ParseResponse(response, err)
}

func (c *OAuthTestClient) PostRevoke(clientBasicAuth, token string) *utils.TestResponse {
	formData := url.Values{}
	formData.Set("token", token)
	encoded := formData.Encode()
	requestUrl := fmt.Sprintf("%s/oauth/revoke", c.baseURL)
	request, _ := http.NewRequest("POST", requestUrl, strings.NewReader(encoded))
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Authorization", fmt.Sprintf("Basic %s", clientBasicAuth))
	response, err := c.httpClient.Do(request)
	return utils.ParseResponse(response, err)
}

func (c *OAuthTestClient) GetMetadata() *utils.TestResponse {
	requestUrl := c.baseURL + "/.well-known/oauth-authorization-server"
	response, err := c.httpClient.Get(requestUrl)
	return utils.ParseResponse(response, err)
}
