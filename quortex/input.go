package quortex

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// GetInput - Get a input
func (c *Client) GetInput(poolName string, inputName string) (*Input, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/inputs/%s", c.HostURL, poolName, inputName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	input := Input{}
	err = json.Unmarshal(body, &input)
	if err != nil {
		return nil, err
	}
	log.Println(input)
	return &input, nil
}

// GetInputs - Returns list of inputs
func (c *Client) GetInputs(poolName string) ([]Input, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/inputs", c.HostURL, poolName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	inputs := []Input{}
	err = json.Unmarshal(body, &inputs)
	if err != nil {
		return nil, err
	}

	return inputs, nil
}

// CreateInput - Create new input
func (c *Client) CreateInput(poolName string, input Input) (*Input, error) {
	log.Println(input)
	rb, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	log.Println(string(rb))
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/pools/%s/inputs", c.HostURL, poolName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)

	if err != nil {
		return nil, err
	}

	newinput := Input{}

	err = json.Unmarshal(body, &newinput)

	log.Println(newinput)

	if err != nil {
		return nil, err
	}
	return &newinput, nil
}

// UpdateInput - Updates a input
func (c *Client) UpdateInput(poolName string, inputName string, input Input) (*Input, error) {

	// First get current streams
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1.0/pools/%s/inputs/%s", c.HostURL, poolName, inputName), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}

	currentInput := Input{}
	err = json.Unmarshal(body, &currentInput)
	if err != nil {
		return nil, err

	}

	// Second create or update streams
	for _, stream := range input.Streams {
		founded := false
		for _, curStream := range currentInput.Streams {
			if stream.Uuid == curStream.Uuid {
				founded = true
			}
		}

		rb, err := json.Marshal(stream)
		if err != nil {
			return nil, err
		}
		if founded == true {
			req, err := http.NewRequest("PUT", fmt.Sprintf("%s/1.0/pools/%s/inputs/%s/streams/%s", c.HostURL, poolName, inputName, stream.Uuid), strings.NewReader(string(rb)))

			if err != nil {
				return nil, err
			}

			_, err = c.doRequest(req)
			if err != nil {
				return nil, err
			}
		} else {
			req, err := http.NewRequest("POST", fmt.Sprintf("%s/1.0/pools/%s/inputs/%s/streams/", c.HostURL, poolName, inputName), strings.NewReader(string(rb)))

			if err != nil {
				return nil, err
			}

			_, err = c.doRequest(req)
			if err != nil {
				return nil, err
			}
		}
	}

	// Third delete streams
	for _, curStream := range currentInput.Streams {
		founded := false
		for _, stream := range input.Streams {
			if stream.Uuid == curStream.Uuid {
				founded = true
			}
		}

		if founded == false {
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/pools/%s/inputs/%s/streams/%s", c.HostURL, poolName, inputName, curStream.Uuid), nil)

			if err != nil {
				return nil, err
			}

			_, err = c.doRequest(req)
			if err != nil {
				return nil, err
			}
		}
	}

	// Finally update input
	rb, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}

	req, err = http.NewRequest("PUT", fmt.Sprintf("%s/1.0/pools/%s/inputs/%s", c.HostURL, poolName, inputName), strings.NewReader(string(rb)))
	if err != nil {
		return nil, err
	}

	body, err = c.doRequest(req)
	if err != nil {
		return nil, err
	}

	updatedinput := Input{}
	err = json.Unmarshal(body, &updatedinput)
	if err != nil {
		return nil, err
	}

	return &updatedinput, nil
}

// DeleteInput - Deletes an input
func (c *Client) DeleteInput(poolName string, inputName string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/1.0/pools/%s/inputs/%s", c.HostURL, poolName, inputName), nil)
	if err != nil {
		return err
	}

	_, err = c.doRequest(req)
	if err != nil {
		return err
	}

	return nil
}
