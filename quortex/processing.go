package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetProcessing - Get a processing
func (c *Client) GetProcessing(poolName string, processingName string) (*Processing, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/processings/%s", c.HostURL, poolName, processingName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	processing := Processing{}
	err = json.Unmarshal(body, &processing)
	if err != nil {
		return nil, err
	}
	return &processing, nil
}

// GetProcessings - Returns list of processings
func (c *Client) GetProcessings(poolName string) ([]Processing, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/processings", c.HostURL, poolName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	processings := []Processing{}
	err = json.Unmarshal(body, &processings)
	if err != nil {
		return nil, err
	}

	return processings, nil
}

// CreateProcessing - Create new processing
func (c *Client) CreateProcessing(poolName string, processing Processing) (*Processing, error) {
	rb, err := json.Marshal(processing)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/pools/%s/processings", c.HostURL, poolName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Publishing-Mode", "explicit")

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newprocessing := Processing{}

	err = json.Unmarshal(body, &newprocessing)

	if err != nil {
		return nil, err
	}
	return &newprocessing, nil
}

// UpdateProcessing - Updates a processing
func (c *Client) UpdateProcessing(poolName string, processingName string, processing Processing) (*Processing, error) {
	rb, err := json.Marshal(processing)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/pools/%s/processings/%s", c.HostURL, poolName, processingName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedprocessing := Processing{}
	err = json.Unmarshal(body, &updatedprocessing)
	if err != nil {
		return nil, err
	}

	return &updatedprocessing, nil
}

// DeleteProcessing - Deletes an processing
func (c *Client) DeleteProcessing(poolName string, processingName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/pools/%s/processings/%s", c.HostURL, poolName, processingName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
