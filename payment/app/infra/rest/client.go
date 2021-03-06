package rest

import (
	"bytes"
	"context"
	"encoding/json"
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
		Url:    baseUrl,
		Client: &http.Client{},
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
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return req, nil
}

func EncodeBody(in interface{}) (io.Reader, error) {
	b := new(bytes.Buffer)

	return b, json.NewEncoder(b).Encode(in)
}

func DecodeBody(res *http.Response, out interface{}) error {
	defer res.Body.Close()
	d := json.NewDecoder(res.Body)

	return d.Decode(out)
}
