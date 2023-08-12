package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) SetProxy(proxy Proxy) *Client {
	proxyURL, err := url.Parse(fmt.Sprintf("http://%s:%s", proxy.Ip, proxy.Port))
	if err != nil {
		c.errorArray = append(c.errorArray, err)
	} else {
		proxyURL.User = url.UserPassword(proxy.UserName, proxy.Password)
		c.httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}
	return c
}

func (c *Client) Request() *Client {
	var err error
	if c.method == http.MethodGet {
		c.httpRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s?%s", c.GetUrl(), c.DataForm.Encode()), nil)
	} else {
		c.httpRequest, err = http.NewRequest(c.method, c.GetUrl(), strings.NewReader(c.DataForm.Encode()))
	}
	c.httpRequest.Header = c.httpHeaders
	c.httpRequest.AddCookie(c.Cookie)
	if err != nil {
		c.errorArray = append(c.errorArray, err)
	}
	return c
}

func (c *Client) Send() HttpResultInterface {
	defer func() {
		if recover() != nil {
			c.errorArray = append(c.errorArray, fmt.Errorf("error: %s", "send request failed"))
			for _, errorInfo := range c.errorArray {
				fmt.Printf("error: %s\n", errorInfo.Error())
			}
		}
	}()
	resp, err := c.httpClient.Do(c.httpRequest)
	if err != nil {
		c.errorArray = append(c.errorArray, err)
	} else {
		return &Response{Body: resp.Body, Resp: resp}
	}
	return nil
}
func (c *Client) NewRequest() HttpResultInterface {
	res := c.Request().Send()
	if res != nil {
		return res
	}
	c.errorArray = append(c.errorArray, fmt.Errorf("error: %s", "send request failed"))
	for _, errorInfo := range c.errorArray {
		fmt.Printf("error: %s\n", errorInfo.Error())
	}
	return HttpResultInterface(nil)
}

func (c *Client) SetCookie(cookie map[string]string) *Client {
	for k, v := range cookie {
		c.Cookie = &http.Cookie{Name: k, Value: v}
	}
	return c
}

func (c *Client) Query(q map[string]interface{}) *Client {
	if q != nil {
		for k, value := range q {
			c.DataForm.Set(k, fmt.Sprintf("%v", value))
		}
	}
	return c
}
func (c *Client) QueryFunc(f func(c *Client)) *Client {
	f(c)
	return c
}

func (c *Client) Headers(h map[string]interface{}) *Client {
	if h != nil {
		for k, value := range h {
			c.httpHeaders.Set(k, fmt.Sprintf("%v", value))
		}
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
	if c.urlSite[len(c.urlSite)-1:] != "/" && c.urlPoint[:1] != "/" {
		c.errorArray = append(c.errorArray, fmt.Errorf("urlSite error: %s%s is not support", c.urlSite, c.urlPoint))
		if c.urlSite[len(c.urlSite)-1:] != "/" {
			c.urlSite = c.urlSite + "/"
		} else if c.urlPoint[:1] != "/" {
			c.urlPoint = "/" + c.urlPoint
		}
	}
	return c.urlSite + c.urlPoint
}
func (c *Client) PostMethod() *Client {
	c.Headers(map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded"})
	return c.Method(http.MethodPost)
}

func (c *Client) GetMethod() *Client {
	c.Headers(map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded"})
	return c.Method(http.MethodGet)
}

func (c *Client) PutMethod() *Client {
	c.Headers(map[string]interface{}{"Content-Type": "application/x-www-form-urlencoded"})
	return c.Method(http.MethodPut)
}

func (c *Client) Method(method string) *Client {
	if strings.Contains(method, strings.Join([]string{http.MethodPost, http.MethodGet, http.MethodPut}, "")) {
		panic("method error: " + method + " is not support")
	}
	c.method = method
	return c
}
