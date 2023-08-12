package requests

import (
	"net/http"
	"net/url"
	"time"
)

func NewClient(timeout int) *Client {
	client := &Client{
		jsonData:    nil,
		Cookie:      &http.Cookie{},
		httpHeaders: http.Header{},
		dataForm:    &url.Values{},
		httpClient:  http.Client{Transport: nil, Timeout: time.Second * time.Duration(timeout)},
	}
	return client
}
