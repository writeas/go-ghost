package admin

import (
	"github.com/writeas/go-ghost"
	"net/http"
)

func AddPost(gc *ghost.Client, p *ghost.PostParams) error {
	c := getC(gc)
	_, err := c.Request(http.MethodPost, c.endpoint("posts"), p)
	return err
}
