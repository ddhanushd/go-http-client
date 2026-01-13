package client

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// HTTPClient is an interface to allow mocking
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// CityClient handles city-related API calls
type CityClient struct {
	client  HTTPClient
	baseURL string
}

// NewCityClient creates a new CityClient
func NewCityClient(client HTTPClient, baseURL string) *CityClient {
	return &CityClient{
		client:  client,
		baseURL: baseURL,
	}
}

// GetCity fetches city data
func (c *CityClient) GetCity(city string) (string, error) {
	url := fmt.Sprintf("%s/cities/%s", c.baseURL, city)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(body)), nil

}
