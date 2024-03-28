package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetOTTEndpoint - Get an ott endpoint
func (c *Client) GetOTTEndpoint(poolName string, ottEndpointName string) (*OTTEndpoint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/ott_endpoints/%s", c.HostURL, poolName, ottEndpointName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	pp := OTTEndpoint{}
	err = json.Unmarshal(body, &pp)
	if err != nil {
		return nil, err
	}
	return &pp, nil
}

// CreateOTTEndpoint - Create new ott endpoint
func (c *Client) CreateOTTEndpoint(poolName string, endpoint OTTEndpoint) (*OTTEndpoint, error) {
	rb, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/pools/%s/ott_endpoints", c.HostURL, poolName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newpp := OTTEndpoint{}

	err = json.Unmarshal(body, &newpp)

	if err != nil {
		return nil, err
	}
	return &newpp, nil
}

// UpdateOTTEndpoint - Updates a ott endpoint
func (c *Client) UpdateOTTEndpoint(poolName string, ottEndpointID string, endpoint OTTEndpoint) (*OTTEndpoint, error) {
	rb, err := json.Marshal(endpoint)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/1.0/pools/%s/ott_endpoints/%s", c.HostURL, poolName, ottEndpointID), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedEndpoint := OTTEndpoint{}
	err = json.Unmarshal(body, &updatedEndpoint)
	if err != nil {
		return nil, err
	}

	return &updatedEndpoint, nil
}

// DeleteOTTEndpoint - Deletes an ott endpoint
func (c *Client) DeleteOTTEndpoint(poolName string, ottEndpointID string) error {
	pp := OTTEndpoint{}
	pp.CustomPath = ""

	rb, err := json.Marshal(pp)
	if err != nil {
		return nil
	}

	log.Println(string(rb))
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/pools/%s/ott_endpoints/%s", c.HostURL, poolName, ottEndpointID), strings.NewReader(string(rb)))
	if err != nil {
		return nil
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil
	}

	updatedpp := OTTEndpoint{}
	err = json.Unmarshal(body, &updatedpp)
	if err != nil {
		return nil
	}

	return nil
}
