package quortex

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

// Client -
type Client struct {
	HostURL      string
	Organization string
	HTTPClient   *http.Client
	Token        string
}

// AuthStruct -
type AuthStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// QxAuthStruct -
type QxAuthStruct struct {
	ApiKey string `json:"api_key_secret"`
}

// QxAuthResponse -
type QxAuthResponse struct {
	Token  string `json:"access_token"`
	Expire string `json:"expires_at"`
}

// OauthStruct -
type OauthStruct struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	GrantType    string `json:"grant_type"`
}

// OauthResponse -
type OauthResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// NewClientApiKey -
func NewClientApiKey(host *string, apikey *string, authserver *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
	}

	// form request body
	rb, err := json.Marshal(QxAuthStruct{
		ApiKey: *apikey,
	})
	if err != nil {
		return nil, err
	}

	// authenticate
	authent := c.HostURL
	if *authserver != "" {
		authent = *authserver
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/token", authent), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	log.Println(req)

	body, err := c.doRequest(req)

	// parse response body
	ar := QxAuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	c.Token = "Bearer " + ar.Token

	return &c, nil
}

// NewClientOauth -
func NewClientOauth(host *string, authserver *string, clientid *string, clientsecret *string, organization *string) (*Client, error) {
	c := Client{
		HTTPClient:   &http.Client{Timeout: 10 * time.Second},
		HostURL:      *host,
		Organization: *organization,
	}

	// form request body
	rb, err := json.Marshal(OauthStruct{
		ClientId:     *clientid,
		ClientSecret: *clientsecret,
		Audience:     fmt.Sprintf("%s/", *host),
		GrantType:    "client_credentials",
	})
	if err != nil {
		return nil, err
	}

	// authenticate
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/oauth/token", *authserver), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	log.Println(req)

	body, err := c.doRequest(req)

	// parse response body
	ar := OauthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, err
	}

	c.Token = "Bearer " + ar.AccessToken

	return &c, nil
}

// NewClientBasicAuth -
func NewClientBasicAuth(host *string, username *string, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
	}

	c.Token = "Basic " + basicAuth(*username, *password)

	return &c, nil
}

// NewClientUnprotected -
func NewClientUnprotected(host *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    *host,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.Token != "" {
		req.Header.Set("Authorization", c.Token)
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated && res.StatusCode != http.StatusNoContent {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
