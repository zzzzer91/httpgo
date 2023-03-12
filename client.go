package httpgo

import (
	"bytes"
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

func (c *Client) Request(method, url string, body []byte, headers ...Header) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "http.NewRequest error")
	}
	for _, f := range headers {
		req.Header.Add(f.Key, f.Val)
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "defaultClient.NewRequest error")
	}
	if resp.StatusCode != http.StatusOK {
		bs, _ := io.ReadAll(resp.Body)
		return nil, errors.Errorf("statusCode is %d, resp: %s", resp.StatusCode, bs)
	}
	return resp, nil
}

func (c *Client) RequestJSON(method, url string, data interface{}, headers ...Header) (*http.Response, error) {
	headers = append(headers, Header{"Content-Type", contentTypeJSON})
	var body []byte
	if data != nil {
		var err error
		body, err = json.Marshal(data)
		if err != nil {
			return nil, errors.Wrap(err, "json.Marshal error")
		}
	}
	return c.Request(method, url, body, headers...)
}

func (c *Client) Get(url string, headers ...Header) (*http.Response, error) {
	return c.Request("GET", url, nil, headers...)
}

func (c *Client) Post(url string, body []byte, headers ...Header) (*http.Response, error) {
	return c.Request("POST", url, body, headers...)
}

func (c *Client) Put(url string, body []byte, headers ...Header) (*http.Response, error) {
	return c.Request("PUT", url, body, headers...)
}

func (c *Client) Delete(url string, headers ...Header) (*http.Response, error) {
	return c.Request("DELETE", url, nil, headers...)
}

func (c *Client) GetJSON(url string, headers ...Header) (*http.Response, error) {
	return c.RequestJSON("GET", url, nil, headers...)
}

func (c *Client) PostJSON(url string, data interface{}, headers ...Header) (*http.Response, error) {
	return c.RequestJSON("POST", url, data, headers...)
}

func (c *Client) PutJSON(url string, data interface{}, headers ...Header) (*http.Response, error) {
	return c.RequestJSON("PUT", url, data, headers...)
}

func (c *Client) DeleteJSON(url string, headers ...Header) (*http.Response, error) {
	return c.RequestJSON("DELETE", url, nil, headers...)
}

func (c *Client) GetJsonWithAuth(url string, token string, headers ...Header) (*http.Response, error) {
	return c.GetJSON(url, append(headers, Header{"Authorization", token})...)
}

func (c *Client) PostJsonWithAuth(url string, data interface{}, token string, headers ...Header) (*http.Response, error) {
	return c.PostJSON(url, data, append(headers, Header{"Authorization", token})...)
}

func (c *Client) PutJsonWithAuth(url string, data interface{}, token string, headers ...Header) (*http.Response, error) {
	return c.PutJSON(url, data, append(headers, Header{"Authorization", token})...)
}

func (c *Client) DeleteJsonWithAuth(url string, token string, headers ...Header) (*http.Response, error) {
	return c.DeleteJSON(url, append(headers, Header{"Authorization", token})...)
}
