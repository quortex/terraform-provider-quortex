package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetTarget - Get a target
func (c *Client) GetTarget(poolName string, targetName string) (*Target, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/targets/%s", c.HostURL, poolName, targetName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	target := Target{}
	err = json.Unmarshal(body, &target)
	if err != nil {
		return nil, err
	}
	return &target, nil
}

// GetTargets - Returns list of targets
func (c *Client) GetTargets(poolName string) ([]Target, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/targets", c.HostURL, poolName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	targets := []Target{}
	err = json.Unmarshal(body, &targets)
	if err != nil {
		return nil, err
	}

	return targets, nil
}

// CreateTarget - Create new target
func (c *Client) CreateTarget(poolName string, target Target) (*Target, error) {
	rb, err := json.Marshal(target)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/pools/%s/targets", c.HostURL, poolName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newtarget := Target{}

	err = json.Unmarshal(body, &newtarget)

	if err != nil {
		return nil, err
	}
	return &newtarget, nil
}

// UpdateTarget - Updates a target
func (c *Client) UpdateTarget(poolName string, targetName string, target Target) (*Target, error) {
	rb, err := json.Marshal(target)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/pools/%s/targets/%s", c.HostURL, poolName, targetName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedtarget := Target{}
	err = json.Unmarshal(body, &updatedtarget)
	if err != nil {
		return nil, err
	}

	return &updatedtarget, nil
}

// DeleteTarget - Deletes an target
func (c *Client) DeleteTarget(poolName string, targetName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/pools/%s/targets/%s", c.HostURL, poolName, targetName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
