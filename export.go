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

func Get(pointPath string, params any) HttpResultInterface {
	return NewClient().GetMethod().UrlSite(pointPath).Query(params).NewRequest()
}

func Post(pointPath string, params any) HttpResultInterface {
	return NewClient().PostMethod().UrlSite(pointPath).Query(params).NewRequest()
}

func Put(pointPath string, params any) HttpResultInterface {
	return NewClient().PutMethod().UrlSite(pointPath).Query(params).NewRequest()
}
