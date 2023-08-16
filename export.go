package requests

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

var DefaultProxyInfo *Proxy
var DefaultCookie []*http.Cookie
var DefaultTimeout time.Duration = 30
var DefaultHeaders http.Header

func NewClient() *Client {
	client := &Client{
		jsonData:    nil,
		Cookie:      DefaultCookie,
		httpHeaders: DefaultHeaders,
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
func Image(imageUrl string, headers map[string]interface{}) ([]byte, error) {
	response := NewClient().GetMethod().UrlSite(imageUrl)
	if headers != nil {
		response.Headers(headers)
	}
	for i := 0; i < 3; i++ {
		result := response.SetRequest().Send()
		if !result.Error() {
			return result.Bytes(), nil
		}
	}
	return nil, fmt.Errorf("failed to get image: %s", imageUrl)
}

func NewImage(imageUrl, imagePath string, headers map[string]interface{}) error {
	if _, ok := os.Stat(imagePath); os.IsNotExist(ok) {
		image, err := Image(imageUrl, headers)
		if err != nil {
			return fmt.Errorf("failed to get image: %w", err)
		}
		err = os.WriteFile(imagePath, image, 0644)
		if err != nil {
			return fmt.Errorf("failed to write image file: %w", err)
		}
	}
	return nil
}
