package client

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io"
	"log"
	"net/http"
	"strings"
)

func (c *Response) Bytes() []byte {
	defer func(Body io.ReadCloser) {
		if c.BodyBytes == nil {
			if ok := Body.Close(); ok != nil {
				log.Printf("Body.Close error: %v", ok)
			}
		}
	}(c.Body)
	if c.BodyBytes != nil {
		return c.BodyBytes
	} else {
		body, err := io.ReadAll(c.Body)
		if err != nil {
			return nil
		}
		c.BodyBytes = body
		return body
	}
}

func (c *Response) Dict() map[string]interface{} {
	var v map[string]interface{}
	err := json.Unmarshal([]byte(c.String()), &v)
	if err != nil {
		log.Printf("json.Unmarshal error: %v", err)
	}
	return v
}
func (c *Response) SetDecodeFunc(f func(c *Response) error) *Response {
	if f != nil {
		if err := f(c); err != nil {
			log.Printf("SetDecodeFunc error: %v", err)
		}
	}
	return c
}
func (c *Response) String() string {
	return string(c.Bytes())
}

func (c *Response) Decode(v any) error {
	return json.NewDecoder(strings.NewReader(c.String())).Decode(v)
}
func (c *Response) DecodePrintError(v any) {
	err := json.NewDecoder(strings.NewReader(c.String())).Decode(v)
	if err != nil {
		log.Printf("json decode error: %v", err)
	}
}
func (c *Response) Json() gjson.Result {
	return gjson.Parse(c.String())
}
func (c *Response) Document() *goquery.Document {
	doc, ok := goquery.NewDocumentFromReader(c.Body)
	if ok != nil {
		log.Printf("goquery.NewDocumentFromReader error: %v", ok)
		return nil
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
