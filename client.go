package httpgo

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

type Client struct {
	c *http.Client
}

type Header struct {
	Key string
	Val string
}

// NewClient create a new httpgo client, if transport is nil, the client uses `http.DefaultTransport`.
func NewClient(timeout time.Duration, transport http.RoundTripper) *Client {
	return &Client{
		c: &http.Client{
			Transport: transport,
			Timeout:   timeout,
		},
	}
}

func (c *Client) Request(ctx context.Context, method, url string, body io.Reader, headers ...Header) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest error")
	}
	for _, f := range headers {
		req.Header.Set(f.Key, f.Val)
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send request error")
	}
	if resp.StatusCode != http.StatusOK {
		bs, _ := io.ReadAll(resp.Body)
		resp.Body.Close()
		return nil, NewStatusError(resp.StatusCode, string(bs))
	}
	return resp, nil
}

func (c *Client) RequestJSON(ctx context.Context, method, url string, data interface{}, headers ...Header) (*http.Response, error) {
	headers = append(headers, Header{"Content-Type", contentTypeJSON})
	var body []byte
	if data != nil {
		var err error
		body, err = json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "json.Marshal error")
		}
	}
	return c.Request(ctx, method, url, bytes.NewReader(body), headers...)
}

func (c *Client) Get(ctx context.Context, url string, headers ...Header) (*http.Response, error) {
	return c.Request(ctx, "GET", url, nil, headers...)
}

func (c *Client) Post(ctx context.Context, url string, body io.Reader, headers ...Header) (*http.Response, error) {
	return c.Request(ctx, "POST", url, body, headers...)
}

func (c *Client) Put(ctx context.Context, url string, body io.Reader, headers ...Header) (*http.Response, error) {
	return c.Request(ctx, "PUT", url, body, headers...)
}

func (c *Client) Delete(ctx context.Context, url string, headers ...Header) (*http.Response, error) {
	return c.Request(ctx, "DELETE", url, nil, headers...)
}

func (c *Client) GetWithAuth(ctx context.Context, url string, token string, headers ...Header) (*http.Response, error) {
	return c.Get(ctx, url, append(headers, Header{"Authorization", token})...)
}

func (c *Client) PostWithAuth(ctx context.Context, url string, token string, body io.Reader, headers ...Header) (*http.Response, error) {
	return c.Post(ctx, url, body, append(headers, Header{"Authorization", token})...)
}

func (c *Client) PutWithAuth(ctx context.Context, url string, token string, body io.Reader, headers ...Header) (*http.Response, error) {
	return c.Put(ctx, url, body, append(headers, Header{"Authorization", token})...)
}

func (c *Client) DeleteWithAuth(ctx context.Context, url string, token string, headers ...Header) (*http.Response, error) {
	return c.Delete(ctx, url, append(headers, Header{"Authorization", token})...)
}

func (c *Client) GetJSON(ctx context.Context, url string, headers ...Header) (*http.Response, error) {
	return c.RequestJSON(ctx, "GET", url, nil, headers...)
}

func (c *Client) PostJSON(ctx context.Context, url string, data interface{}, headers ...Header) (*http.Response, error) {
	return c.RequestJSON(ctx, "POST", url, data, headers...)
}

func (c *Client) PutJSON(ctx context.Context, url string, data interface{}, headers ...Header) (*http.Response, error) {
	return c.RequestJSON(ctx, "PUT", url, data, headers...)
}

func (c *Client) DeleteJSON(ctx context.Context, url string, headers ...Header) (*http.Response, error) {
	return c.RequestJSON(ctx, "DELETE", url, nil, headers...)
}

func (c *Client) GetJsonWithAuth(ctx context.Context, url string, token string, headers ...Header) (*http.Response, error) {
	return c.GetJSON(ctx, url, append(headers, Header{"Authorization", token})...)
}

func (c *Client) PostJsonWithAuth(ctx context.Context, url string, token string, data interface{}, headers ...Header) (*http.Response, error) {
	return c.PostJSON(ctx, url, data, append(headers, Header{"Authorization", token})...)
}

func (c *Client) PutJsonWithAuth(ctx context.Context, url string, token string, data interface{}, headers ...Header) (*http.Response, error) {
	return c.PutJSON(ctx, url, data, append(headers, Header{"Authorization", token})...)
}

func (c *Client) DeleteJsonWithAuth(ctx context.Context, url string, token string, headers ...Header) (*http.Response, error) {
	return c.DeleteJSON(ctx, url, append(headers, Header{"Authorization", token})...)
}
