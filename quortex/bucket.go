package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetBucket - Get a bucket
func (c *Client) GetBucket(dataplaneName string, bucketName string) (*Bucket, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/private/dataplanes/%s/buckets/%s", c.HostURL, dataplaneName, bucketName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	bucket := Bucket{}
	err = json.Unmarshal(body, &bucket)
	if err != nil {
		return nil, err
	}
	log.Println(bucket)
	return &bucket, nil
}

// GetBuckets - Returns list of buckets
func (c *Client) GetBuckets(dataplaneName string) ([]Bucket, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/private/dataplanes/%s/buckets", c.HostURL, dataplaneName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	buckets := []Bucket{}
	err = json.Unmarshal(body, &buckets)
	if err != nil {
		return nil, err
	}

	return buckets, nil
}

// CreateBucket - Create new bucket
func (c *Client) CreateBucket(dataplaneName string, bucket Bucket) (*Bucket, error) {
	log.Println(bucket)
	rb, err := json.Marshal(bucket)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/private/dataplanes/%s/buckets", c.HostURL, dataplaneName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newbucket := Bucket{}

	err = json.Unmarshal(body, &newbucket)

	log.Println(newbucket)

	if err != nil {
		return nil, err
	}
	return &newbucket, nil
}

// UpdateBucket - Updates a bucket
func (c *Client) UpdateBucket(dataplaneName string, bucketName string, bucket Bucket) (*Bucket, error) {

	rb, err := json.Marshal(bucket)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/private/dataplanes/%s/buckets/%s", c.HostURL, dataplaneName, bucketName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedbucket := Bucket{}
	err = json.Unmarshal(body, &updatedbucket)
	if err != nil {
		return nil, err
	}

	return &updatedbucket, nil
}

// DeleteBucket - Deletes an bucket
func (c *Client) DeleteBucket(dataplaneName string, bucketName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/private/dataplanes/%s/buckets/%s", c.HostURL, dataplaneName, bucketName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
