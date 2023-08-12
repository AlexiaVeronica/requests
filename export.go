package client

import (
	"net/http"
	"net/url"
)

func NewClient(data interface{}) *Client {
	client := &Client{Cookie: &http.Cookie{}, httpHeaders: http.Header{}, httpClient: http.Client{Transport: nil}}
	if dataParams := data.(url.Values); dataParams != nil {
		var params = url.Values{}
		for k, v := range dataParams {
			params.Set(k, v[0])
		}
		client.DataForm = &params
		return client
	}
	client.DataForm = &url.Values{}
	return client
}
