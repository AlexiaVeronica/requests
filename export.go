package requests

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

var defaultProxyInfo *Proxy

func SetDefaultProxy(ip string, port string, username string, psw string) {
	defaultProxyInfo = &Proxy{Ip: ip, Port: port, UserName: username, Password: psw}
}

var defaultCookie map[string]string

func SetDefaultCookie(cookie map[string]string) {
	defaultCookie = cookie
}

var defaultTimeout time.Duration

func SetDefaultTimeout(timeout int) {
	defaultTimeout = time.Duration(timeout)
}

var defaultHeaders http.Header

func SetDefaultHeaders(headers map[string]interface{}) {
	defaultHeaders = make(http.Header)
	for k, v := range headers {
		defaultHeaders.Set(k, fmt.Sprintf("%v", v))
	}
}

func NewClient() *Client {
	client := &Client{
		jsonData:    nil,
		Cookie:      make(map[string]string),
		httpHeaders: make(http.Header),
		dataForm:    &url.Values{},
		httpClient:  http.Client{Timeout: time.Second * 30},
	}
	if defaultTimeout != 0 {
		client.httpClient.Timeout = time.Second * defaultTimeout
	}

	if defaultHeaders != nil {
		client.httpHeaders = defaultHeaders
	}

	if defaultCookie != nil {
		client.SetCookie(defaultCookie)
	}

	if defaultProxyInfo != nil {
		SetProxy(client, defaultProxyInfo)
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
