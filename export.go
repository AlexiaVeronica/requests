package requests

import (
	"net/http"
	"net/url"
	"time"
)

var DefaultProxyInfo *Proxy
var DefaultTimeout time.Duration = 30

func NewClient() *Client {
	client := &Client{
		jsonData:    nil,
		Cookie:      &http.Cookie{},
		httpHeaders: http.Header{},
		dataForm:    &url.Values{},
		httpClient:  http.Client{Transport: nil, Timeout: time.Second * DefaultTimeout},
	}
	if DefaultProxyInfo != nil {
		return client.SetProxy(DefaultProxyInfo)
	}
	return client
}

func Get(pointPath string, params any, headers map[string]interface{}) HttpResultInterface {
	response := NewClient().GetMethod().UrlSite(pointPath).Query(params)
	if headers != nil {
		response.Headers(headers)
	}
	for i := 0; i < 3; i++ {
		result := response.NewRequest()
		if result != nil {
			return result
		}
	}
	return nil
}

func Post(pointPath string, params any, headers map[string]interface{}) HttpResultInterface {
	response := NewClient().PostMethod().UrlSite(pointPath).Query(params)
	if headers != nil {
		response.Headers(headers)
	}
	for i := 0; i < 3; i++ {
		result := response.NewRequest()
		if result != nil {
			return result
		}
	}
	return nil
}

func Put(pointPath string, params any, headers map[string]interface{}) HttpResultInterface {
	response := NewClient().PutMethod().UrlSite(pointPath).Query(params)
	if headers != nil {
		response.Headers(headers)
	}
	for i := 0; i < 3; i++ {
		result := response.NewRequest()
		if result != nil {
			return result
		}
	}
	return nil
}
