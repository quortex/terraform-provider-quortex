package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetWebhook - Get a webhook
func (c *Client) GetWebhook(webhookName string) (*Webhook, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/organization/webhooks/%s", c.HostURL, webhookName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	webhook := Webhook{}
	err = json.Unmarshal(body, &webhook)
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

// GetWebhooks - Returns list of webhooks
func (c *Client) GetWebhooks() ([]Webhook, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/organization/webhooks", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	webhooks := []Webhook{}
	err = json.Unmarshal(body, &webhooks)
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

// CreateWebhook - Create new webhook
func (c *Client) CreateWebhook(webhook Webhook) (*Webhook, error) {
	rb, err := json.Marshal(webhook)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] %s !", rb)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/organization/webhooks", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newwebhook := Webhook{}
	err = json.Unmarshal(body, &newwebhook)

	if err != nil {

		return nil, err
	}

	return &newwebhook, nil
}

// UpdateWebhook - Updates a webhook
func (c *Client) UpdateWebhook(webhookName string, webhook Webhook) (*Webhook, error) {
	rb, err := json.Marshal(webhook)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/organization/webhooks/%s", c.HostURL, webhookName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedwebhook := Webhook{}
	err = json.Unmarshal(body, &updatedwebhook)
	if err != nil {
		return nil, err
	}

	return &updatedwebhook, nil
}

// DeleteWebhook - Deletes an webhook
func (c *Client) DeleteWebhook(webhookName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/organization/webhooks/%s", c.HostURL, webhookName, webhookName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
