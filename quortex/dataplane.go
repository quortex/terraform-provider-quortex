package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetDataplane - Get a dataplane
func (c *Client) GetDataplane(dataplaneName string) (*Dataplane, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/private/dataplanes/%s?%s", c.HostURL, dataplaneName, c.Organization), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dataplane := Dataplane{}
	err = json.Unmarshal(body, &dataplane)
	if err != nil {
		return nil, err
	}

	return &dataplane, nil
}

// GetDataplanes - Returns list of dataplanes
func (c *Client) GetDataplanes() ([]Dataplane, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/private/dataplanes?%s", c.HostURL, c.Organization), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	dataplanes := []Dataplane{}
	err = json.Unmarshal(body, &dataplanes)
	if err != nil {
		return nil, err
	}

	return dataplanes, nil
}

// CreateDataplane - Create new dataplane
func (c *Client) CreateDataplane(dataplane Dataplane) (*Dataplane, error) {
	rb, err := json.Marshal(dataplane)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] %s !", rb)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/private/dataplanes?%s", c.HostURL, c.Organization), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newdataplane := Dataplane{}
	err = json.Unmarshal(body, &newdataplane)

	if err != nil {

		return nil, err
	}

	return &newdataplane, nil
}

// UpdateDataplane - Updates a dataplane
func (c *Client) UpdateDataplane(dataplaneName string, dataplane Dataplane) (*Dataplane, error) {
	rb, err := json.Marshal(dataplane)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/private/dataplanes/%s?%s", c.HostURL, dataplaneName, c.Organization), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateddataplane := Dataplane{}
	err = json.Unmarshal(body, &updateddataplane)
	if err != nil {
		return nil, err
	}

	return &updateddataplane, nil
}

// DeleteDataplane - Deletes an dataplane
func (c *Client) DeleteDataplane(dataplaneName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/private/dataplanes/%s?dataplane_uuid=%s&%s", c.HostURL, dataplaneName, dataplaneName, c.Organization), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
