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
		SetProxy(client, DefaultProxyInfo)
	}
	return client
}

func Get(pointPath string, params any, headers map[string]interface{}) HttpResultInterface {
	response := NewClient().GetMethod().UrlSite(pointPath).FormQuery(params)
	if headers != nil {
		response.Headers(headers)
	}
	for i := 0; i < 3; i++ {
		result := response.SetRequest().Send()
		if !result.Error() {
			return result
		}
	}
	return NewResponse(nil)
}

func Post(pointPath string, params any, headers map[string]interface{}) HttpResultInterface {
	response := NewClient().PostMethod().UrlSite(pointPath)
	if headers["Content-Type"] == ContentTypeForm {
		response.FormQuery(params)
	} else if headers["Content-Type"] == ContentTypeJson {
		response.JsonQuery(params)
	}
	if headers != nil {
		response.Headers(headers)
	}
	for i := 0; i < 3; i++ {
		result := response.SetRequest().Send()
		if !result.Error() {
			return result
		}
	}
	return NewResponse(nil)
}

func Put(pointPath string, params any, headers map[string]interface{}) HttpResultInterface {
	response := NewClient().PutMethod().UrlSite(pointPath)
	if headers["Content-Type"] == ContentTypeForm {
		response.FormQuery(params)
	} else if headers["Content-Type"] == ContentTypeJson {
		response.JsonQuery(params)
	}
	if headers != nil {
		response.Headers(headers)
	}
	for i := 0; i < 3; i++ {
		result := response.SetRequest().Send()
		if !result.Error() {
			return result
		}
	}
	return NewResponse(nil)
}
