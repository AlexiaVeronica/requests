package requests

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	method      string
	urlPoint    string
	urlSite     string
	jsonData    []byte
	dataForm    *url.Values
	httpHeaders http.Header
	httpRequest *http.Request
	httpClient  http.Client
	Cookie      *http.Cookie
}

type HttpClientInterface interface {
	Query(data interface{}, queryType string) *Client
	JsonQuery(data interface{}) *Client
	FormQuery(data interface{}) *Client
	QueryFunc(f func(c *Client) (interface{}, string)) *Client
	QueryResult() io.Reader
	Headers(m map[string]interface{}) *Client
	Header(k string, value interface{}) *Client
	HeadersFunc(f func(c *Client)) *Client
	SetCookie(cookie map[string]string) *Client
	Method(method string) *Client
	GetMethod() *Client
	PostMethod() *Client
	PutMethod() *Client
	UrlPoint(urlPoint string) *Client
	GetUrl() string
	UrlSite(urlSite string) *Client

	SetRequest() HttpRequestInterface
}

func (c *Client) SetRequest() HttpRequestInterface {
	var err error
	if c.method == "" {
		c.method = http.MethodGet
	}
	if c.method == http.MethodGet {
		c.httpRequest, err = http.NewRequest(http.MethodGet, c.GetUrl(), nil)
		c.httpRequest.URL.RawQuery = c.dataForm.Encode()
	} else {
		c.httpRequest, err = http.NewRequest(c.method, c.GetUrl(), c.QueryResult())
	}
	c.httpRequest.Header = c.httpHeaders
	c.httpRequest.AddCookie(c.Cookie)
	if err != nil {
		log.Printf("http.NewRequest error: %v", err)
		return &Request{httpRequest: nil}
	}
	return &Request{httpRequest: c.httpRequest, httpClient: c.httpClient, url: c.GetUrl()}
}

func (c *Client) SetCookie(cookie map[string]string) *Client {
	for k, v := range cookie {
		c.Cookie = &http.Cookie{Name: k, Value: v}
	}
	return c
}

func (c *Client) Header(k string, value interface{}) *Client {
	c.httpHeaders.Set(k, fmt.Sprintf("%v", value))
	return c
}

func (c *Client) Headers(m map[string]interface{}) *Client {
	for k, v := range m {
		c.Header(k, v)
	}
	return c
}
func (c *Client) HeadersFunc(f func(c *Client)) *Client {
	f(c)
	return c
}
func (c *Client) UrlSite(urlSite string) *Client {
	if !strings.Contains(urlSite, "http") {
		panic("urlSite error: " + urlSite + " is not support")
	}
	c.urlSite = urlSite
	return c
}

func (c *Client) UrlPoint(urlPoint string) *Client {
	c.urlPoint = urlPoint
	return c
}
func (c *Client) GetUrl() string {
	if strings.TrimSpace(c.urlPoint) != "" {
		if c.urlSite[len(c.urlSite)-1:] != "/" && c.urlPoint[:1] != "/" {
			log.Printf("urlSite error: %s%s is not support", c.urlSite, c.urlPoint)
			c.urlPoint = "/" + c.urlPoint
		}
	}
	return c.urlSite + c.urlPoint
}
func (c *Client) PostMethod() *Client {
	return c.Method(http.MethodPost)
}

func (c *Client) GetMethod() *Client {
	return c.Method(http.MethodGet)
}

func (c *Client) PutMethod() *Client {
	return c.Method(http.MethodPut)
}

func (c *Client) Method(method string) *Client {
	if strings.Contains(method, strings.Join([]string{http.MethodPost, http.MethodGet, http.MethodPut}, "")) {
		panic("method error: " + method + " is not support")
	}
	c.method = method
	return c
}
