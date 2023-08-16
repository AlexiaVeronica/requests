package requests

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Err       error
	BodyBytes []byte
	Body      io.ReadCloser
	Resp      *http.Response
}

type HttpResultInterface interface {
	SetDecodeFunc(func(f *Response) error) *Response
	Bytes() []byte
	Error() bool
	String() string
	Json() gjson.Result
	Map() map[string]interface{}
	Decode(v any) error
	GetCookie() []*http.Cookie
	GetHeader() http.Header
	GetStatusCode() int
	Document() *goquery.Document
	GetStatus() string
	GetQuery() string
	GetUrl() string
}

func NewResponse(resp *http.Response) HttpResultInterface {
	res := &Response{}
	if resp == nil {
		res.Err = fmt.Errorf("response is nil, please check your network")
		return res
	} else if resp.Body == nil {
		res.Err = fmt.Errorf("response body is nil, please check your network")
		return res
	}

	res.Resp = resp
	res.Body = resp.Body
	return res
}
func (c *Response) Error() bool {
	if c.Err != nil {
		log.Printf("Response error: %v", c.Err)
		return true
	}
	return false
}
func (c *Response) Bytes() []byte {
	if c.Err != nil {
		log.Printf("Response error: %v", c.Err)
		return nil
	}
	defer func(Body io.ReadCloser) {
		if c.BodyBytes == nil {
			if ok := Body.Close(); ok != nil {
				log.Printf("Body.Close error: %v", ok)
			}
		}
	}(c.Body)
	// the body has been read, so we don't need to read it again
	if c.BodyBytes != nil {
		return c.BodyBytes
	}

	body, err := io.ReadAll(c.Body)
	if err != nil {
		log.Printf("io.ReadAll error: %v", err)
		return nil
	}
	c.BodyBytes = body
	return body

}

func (c *Response) String() string {
	return string(c.Bytes())
}

func (c *Response) Map() map[string]interface{} {
	var v map[string]interface{}
	err := json.Unmarshal([]byte(c.String()), &v)
	if err != nil {
		log.Printf("json.Unmarshal error: %v", err)
	}
	return v
}

func (c *Response) SetDecodeFunc(f func(c *Response) error) *Response {
	if f != nil {
		if c.Err != nil {
			log.Printf("SetDecodeFunc error: %v", c.Err)
			return c
		}

		if err := f(c); err != nil {
			log.Printf("SetDecodeFunc error: %v", err)
		}
	} else {
		log.Printf("SetDecodeFunc error: %v", "f is nil")
	}
	return c
}

func (c *Response) Decode(v any) error {
	if responseString := c.String(); responseString == "" {
		return fmt.Errorf("response string is nil")
	} else {
		return json.NewDecoder(strings.NewReader(responseString)).Decode(v)
	}
}

func (c *Response) Json() gjson.Result {
	return gjson.Parse(c.String())
}

func (c *Response) Document() *goquery.Document {
	doc, ok := goquery.NewDocumentFromReader(c.Body)
	if ok != nil {
		log.Printf("goquery.NewDocumentFromReader error: %v", ok)
	}
	return doc
}

func (c *Response) GetCookie() []*http.Cookie {
	return c.Resp.Cookies()
}

func (c *Response) GetHeader() http.Header {
	return c.Resp.Header
}

func (c *Response) GetStatusCode() int {
	return c.Resp.StatusCode
}

func (c *Response) GetStatus() string {
	return c.Resp.Status
}

func (c *Response) GetQuery() string {
	return c.Resp.Request.URL.RawQuery
}

func (c *Response) GetUrl() string {
	return c.Resp.Request.URL.String()
}
