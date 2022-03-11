package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetPool - Get a pool
func (c *Client) GetPool(poolName string) (*Pool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s", c.HostURL, poolName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	pool := Pool{}
	err = json.Unmarshal(body, &pool)
	if err != nil {
		return nil, err
	}

	return &pool, nil
}

// GetPools - Returns list of pools
func (c *Client) GetPools() ([]Pool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	pools := []Pool{}
	err = json.Unmarshal(body, &pools)
	if err != nil {
		return nil, err
	}

	return pools, nil
}

// CreatePool - Create new pool
func (c *Client) CreatePool(pool Pool) (*Pool, error) {
	rb, err := json.Marshal(pool)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] %s !", rb)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/pools", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newpool := Pool{}
	err = json.Unmarshal(body, &newpool)

	if err != nil {

		return nil, err
	}

	return &newpool, nil
}

// UpdatePool - Updates a pool
func (c *Client) UpdatePool(poolName string, pool Pool) (*Pool, error) {
	rb, err := json.Marshal(pool)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/pools/%s", c.HostURL, poolName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedpool := Pool{}
	err = json.Unmarshal(body, &updatedpool)
	if err != nil {
		return nil, err
	}

	return &updatedpool, nil
}

// DeletePool - Deletes an pool
func (c *Client) DeletePool(poolName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/pools/%s?pool_uuid=%s", c.HostURL, poolName, poolName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
