// Package admin supports interacting with the Ghost Admin API.
package admin

import (
	"fmt"
	"github.com/writeas/go-ghost"
)

type Client struct {
	ghost.Client
}

func getC(c *ghost.Client) *Client {
	return &Client{*c}
}

func (c *Client) endpoint(resource string) string {
	return fmt.Sprintf("/%s/api/%s/admin/%s/", c.GhostPath, c.Version, resource)
}

func (c *Client) endpointForID(resource, id string) string {
	return fmt.Sprintf("/%s/api/%s/admin/%s/%s/", c.GhostPath, c.Version, resource, id)
}

func (c *Client) endpointForSlug(resource, slug string) string {
	return fmt.Sprintf("/%s/api/%s/admin/%s/slug/%s/", c.GhostPath, c.Version, resource, slug)
}
