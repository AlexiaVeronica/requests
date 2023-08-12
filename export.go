package client

import (
	"net/http"
	"net/url"
)

func NewClient() *Client {
	client := &Client{
		Cookie:      &http.Cookie{},
		httpHeaders: http.Header{},
		httpClient:  http.Client{Transport: nil},
		dataForm:    &url.Values{},
		jsonData:    nil,
	}
	return client
}
