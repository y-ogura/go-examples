// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "file": file Resource Client
//
// Command:
// $ goagen
// --design=github.com/go-examples/goa/response_file/design
// --out=$(GOPATH)/src/github.com/go-examples/goa/response_file
// --version=v1.2.0-dirty

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

// FileFilePath computes a request path to the file action of file.
func FileFilePath() string {

	return fmt.Sprintf("/downloads/")
}

// Download file
func (c *Client) FileFile(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewFileFileRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewFileFileRequest create the request corresponding to the file action endpoint of the file resource.
func (c *Client) NewFileFileRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
