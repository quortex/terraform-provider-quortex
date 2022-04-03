package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetDrmMerchant - Get a drmmerchant
func (c *Client) GetDrmMerchant(drmmerchantName string) (*DrmMerchant, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/drm_merchants/%s", c.HostURL, drmmerchantName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	drmmerchant := DrmMerchant{}
	err = json.Unmarshal(body, &drmmerchant)
	if err != nil {
		return nil, err
	}

	return &drmmerchant, nil
}

// GetDrmMerchants - Returns list of drmmerchants
func (c *Client) GetDrmMerchants() ([]DrmMerchant, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/drm_merchants", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	drmmerchants := []DrmMerchant{}
	err = json.Unmarshal(body, &drmmerchants)
	if err != nil {
		return nil, err
	}

	return drmmerchants, nil
}

// CreateDrmMerchant - Create new drmmerchant
func (c *Client) CreateDrmMerchant(drmmerchant DrmMerchant) (*DrmMerchant, error) {
	rb, err := json.Marshal(drmmerchant)
	if err != nil {
		return nil, err
	}
	log.Printf("[INFO] %s !", rb)

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/drm_merchants", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newdrmmerchant := DrmMerchant{}
	err = json.Unmarshal(body, &newdrmmerchant)

	if err != nil {

		return nil, err
	}

	return &newdrmmerchant, nil
}

// UpdateDrmMerchant - Updates a drmmerchant
func (c *Client) UpdateDrmMerchant(drmmerchantName string, drmmerchant DrmMerchant) (*DrmMerchant, error) {
	rb, err := json.Marshal(drmmerchant)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/drm_merchants/%s", c.HostURL, drmmerchantName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updateddrmmerchant := DrmMerchant{}
	err = json.Unmarshal(body, &updateddrmmerchant)
	if err != nil {
		return nil, err
	}

	return &updateddrmmerchant, nil
}

// DeleteDrmMerchant - Deletes an drmmerchant
func (c *Client) DeleteDrmMerchant(drmmerchantName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/drm_merchants/%s?drm_merchant_uuid=%s", c.HostURL, drmmerchantName, drmmerchantName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
