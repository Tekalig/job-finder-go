// hasura/client.go
package hasura

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Client struct {
	endpoint    string
	adminSecret string
	httpClient  *http.Client
}

func NewClient(endpoint, adminSecret string) *Client {
	return &Client{
		endpoint:    endpoint,
		adminSecret: adminSecret,
		httpClient:  &http.Client{},
	}
}

func (c *Client) Execute(query string, variables map[string]interface{}, result interface{}) error {
	payload := map[string]interface{}{
		"query":     query,
		"variables": variables,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", c.endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-hasura-admin-secret", c.adminSecret)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}
