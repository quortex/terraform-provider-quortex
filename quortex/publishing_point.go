package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetPublishingPoint - Get a publishing point
func (c *Client) GetPublishingPoint(poolName string, publishingPointName string) (*PublishingPoint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/publishing_points/%s", c.HostURL, poolName, publishingPointName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	pp := PublishingPoint{}
	err = json.Unmarshal(body, &pp)
	if err != nil {
		return nil, err
	}
	return &pp, nil
}

// GetPublishingPoints - Returns list of publishing points
func (c *Client) GetPublishingPoints(poolName string) ([]PublishingPoint, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/publishing_points", c.HostURL, poolName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	pps := []PublishingPoint{}
	err = json.Unmarshal(body, &pps)
	if err != nil {
		return nil, err
	}

	return pps, nil
}

// CreatePublishingPoint - Create new publishing point
func (c *Client) CreatePublishingPoint(poolName string, pp PublishingPoint) (*PublishingPoint, error) {

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/publishing_points", c.HostURL, poolName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	pps := []PublishingPoint{}
	err = json.Unmarshal(body, &pps)
	if err != nil {
		return nil, err
	}

	publishingPointName := ""
	for _, ppp := range pps {
		if (pp.InputUuid == ppp.InputUuid) && (pp.ProcessingUuid == ppp.ProcessingUuid) && (pp.TargetUuid == ppp.TargetUuid) {
			publishingPointName = ppp.Uuid
			break
		}
	}

	rb, err := json.Marshal(pp)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err = http.NewRequest("PUT", fmt.Sprintf("%s/1.0/pools/%s/publishing_points/%s", c.HostURL, poolName, publishingPointName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err = c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newpp := PublishingPoint{}

	err = json.Unmarshal(body, &newpp)

	if err != nil {
		return nil, err
	}
	return &newpp, nil
}

// UpdatePublishingPoint - Updates a publishing point
func (c *Client) UpdatePublishingPoint(poolName string, publishingPointName string, pp PublishingPoint) (*PublishingPoint, error) {
	rb, err := json.Marshal(pp)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/pools/%s/publishing_points/%s", c.HostURL, poolName, publishingPointName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedpp := PublishingPoint{}
	err = json.Unmarshal(body, &updatedpp)
	if err != nil {
		return nil, err
	}

	return &updatedpp, nil
}

// DeletePublishingPoint - Deletes an publishing point
func (c *Client) DeletePublishingPoint(poolName string, publishingPointName string) error {
	pp := PublishingPoint{}
	pp.Published = true
	pp.CustomPath = ""

	rb, err := json.Marshal(pp)
	if err != nil {
		return nil
	}

	log.Println(string(rb))
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/1.0/pools/%s/publishing_points/%s", c.HostURL, poolName, publishingPointName), strings.NewReader(string(rb)))
	if err != nil {
		return nil
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil
	}

	updatedpp := PublishingPoint{}
	err = json.Unmarshal(body, &updatedpp)
	if err != nil {
		return nil
	}

	return nil
}
