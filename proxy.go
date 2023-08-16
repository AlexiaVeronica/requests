package requests

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"reflect"
)

type Proxy struct {
	Ip       string
	Port     string
	UserName string
	Password string
}

func SetProxy(c *Client, proxy *Proxy) {
	if proxy == nil {
		log.Println("proxy is nil")
	} else {
		value := reflect.ValueOf(proxy).Elem()
		for i := 0; i < value.NumField(); i++ {
			if value.Field(i).String() == "" {
				log.Printf("error: %s", "proxy field is empty")
				return
			}
		}
		proxyURL, err := url.Parse(fmt.Sprintf("http://%s:%s", proxy.Ip, proxy.Port))
		if err != nil {
			log.Printf("url.Parse error: %v", err)
		} else {
			proxyURL.User = url.UserPassword(proxy.UserName, proxy.Password)
			c.httpClient.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
		}
	}
}
