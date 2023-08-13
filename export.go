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

func Get(pointPath string, params any) HttpResultInterface {
	return NewClient(30).GetMethod().UrlSite(pointPath).Query(params).NewRequest()
}

func Post(pointPath string, params any) HttpResultInterface {
	return NewClient(30).PostMethod().UrlSite(pointPath).Query(params).NewRequest()
}

func Put(pointPath string, params any) HttpResultInterface {
	return NewClient(30).PutMethod().UrlSite(pointPath).Query(params).NewRequest()
}
