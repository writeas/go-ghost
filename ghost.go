package ghost

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io"
	"net/http"
	"time"
)

type Client struct {
	URL       string
	Key       string
	Version   string
	GhostPath string
	UserAgent string
	client    *http.Client
}

// defaultHTTPTimeout is the default http.Client timeout.
const defaultHTTPTimeout = 10 * time.Second

// NewClient creates a new API client.
func NewClient(url, key string) *Client {
	httpClient := &http.Client{Timeout: defaultHTTPTimeout}
	return &Client{
		URL:       url,
		Key:       key,
		Version:   "v2",
		GhostPath: "ghost",
		client:    httpClient,
	}
}

func (c *Client) Request(method, path string, data interface{}) (map[string][]interface{}, error) {
	var b *bytes.Buffer
	if method == http.MethodPost {
		b = &bytes.Buffer{}
		json.NewEncoder(b).Encode(data)
	}
	r, err := c.buildRequest(method, path, b)
	if err != nil {
		return nil, err
	}

	res, err := c.doRequest(r)
	// Check for transport errors
	if err != nil {
		return nil, err
	}
	// Check for application errors
	if em, ok := res["errors"]; ok {
		e := Error{}
		err = mapstructure.Decode(em[0], &e)
		if err != nil {
			return nil, err
		}
		return nil, e
	}
	return res, nil
}

func (c *Client) buildRequest(method, path string, data io.Reader) (*http.Request, error) {
	url := fmt.Sprintf("%s%s", c.URL, path)
	r, err := http.NewRequest(method, url, data)
	if err != nil {
		return nil, fmt.Errorf("buildRequest: %v", err)
	}
	c.prepareRequest(r)

	return r, nil
}

func (c *Client) doRequest(r *http.Request) (map[string][]interface{}, error) {
	resp, err := c.client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("doRequest: %v", err)
	}
	defer resp.Body.Close()

	env := map[string][]interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&env)
	if err != nil {
		return nil, err
	}

	return env, nil
}

func (c *Client) prepareRequest(r *http.Request) {
	ua := c.UserAgent
	if ua == "" {
		ua = "go-ghost v1"
	}
	r.Header.Add("User-Agent", ua)
	r.Header.Add("Content-Type", "application/json")
}

func (c *Client) endpointForID(api, resource, id string) string {
	return fmt.Sprintf("/%s/api/%s/%s/%s/%s/", c.GhostPath, c.Version, api, resource, id)
}

func (c *Client) endpointForSlug(api, resource, slug string) string {
	return fmt.Sprintf("/%s/api/%s/%s/%s/slug/%s/", c.GhostPath, c.Version, api, resource, slug)
}

// String returns a pointer to the string value passed in.
func String(v string) *string {
	return &v
}

// Bool returns a pointer to the bool value passed in.
func Bool(v bool) *bool {
	return &v
}
