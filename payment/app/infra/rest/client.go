package rest

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"path"
)

type Client struct {
	Url *url.URL
	*http.Client
}

func NewHttpClient(baseUrl *url.URL) *Client {
	return &Client{
		Url: baseUrl,
	}
}

func (c *Client) NewReq(ctx context.Context, method, spath string, body io.Reader) (*http.Request, error) {
	u := c.Url
	u.Path = path.Join(c.Url.Path, spath)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}
