package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const BaseURL = "https://api.shodan.io"

type Client struct {
	apiKey string
}

func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

func (s *Client) APIInfo() (*APIInfo, error) {
	reqURL := fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey)
	res, err := http.Get(reqURL)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (s *Client) HostSearch() {
	return
}
