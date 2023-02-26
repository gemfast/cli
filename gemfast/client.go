package gemfast

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"
)


type Client struct {
	BaseUrl string
	AuthMethod AuthMethod
}

func NewClient(url string, auth AuthMethod) (*Client) {
	return &Client{BaseUrl: url, AuthMethod: auth}
}

func (c *Client) sendRequest(action string, path string, body []byte) (*http.Response, error) {
	var bodyReader io.Reader
	if body != nil {
		bodyReader = bytes.NewBuffer(body)
	}
	req, err := http.NewRequest(strings.ToUpper(action), c.BaseUrl + path, bodyReader)
	if err != nil {
		return nil, err
	}

	c.applyHeaders(req)

	return http.DefaultClient.Do(req)
}

func (c *Client) applyHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if c.AuthMethod != nil {
		switch c.AuthMethod.(type) {
		case LocalAuth:
			req.Header.Set("Authorization", c.AuthMethod.String())
		default:
		  fmt.Errorf("unknown auth method %s", c.AuthMethod)
		}		
	}

	return
}

func (c *Client) head(path string) (*http.Response, error) {
	return c.sendRequest("HEAD", path, nil)
}

func (c *Client) get(path string) (*http.Response, error) {
	return c.sendRequest("GET", path, nil)
}

func (c *Client) post(path string, body []byte) (*http.Response, error) {
	return c.sendRequest("POST", path, body)
}

func (c *Client) put(path string, body []byte) (*http.Response, error) {
	return c.sendRequest("PUT", path, body)
}

func (c *Client) delete(path string, body []byte) (*http.Response, error) {
	return c.sendRequest("DELETE", path, body)
}