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

// HostURL - Default Hashicups URL
const HostURL string = "https://api.quortex.io"

// Client -
type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
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

// AuthResponse -
type AuthResponse struct {
	UserID   int    `json:"user_id`
	Username string `json:"username`
	Token    string `json:"token"`
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

// NewClient -
func NewClient(host, username *string, password *string, apikey *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		// Default Hashicups URL
		HostURL: HostURL,
	}

	if host != nil {
		c.HostURL = *host
	}

	if apikey != nil {
		// form request body
		rb, err := json.Marshal(QxAuthStruct{
			ApiKey: *apikey,
		})
		if err != nil {
			return nil, err
		}

		// authenticate
		req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/token", c.HostURL), strings.NewReader(string(rb)))
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
	}

	if (username != nil) && (password != nil) {
		c.Token = "Basic " + basicAuth(*username, *password)
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
