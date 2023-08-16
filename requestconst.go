package requests

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
)

const (
	ContentTypeJson       = "application/json"
	ContentTypeForm       = "application/x-www-form-urlencoded"
	ContentTypeFile       = "multipart/form-data"
	ContentTypeFormString = "json"
	ContentTypeJsonString = "form"
)

type Proxy struct {
	Ip       string
	Port     string
	UserName string
	Password string
}

type Response struct {
	Err       error
	BodyBytes []byte
	Body      io.ReadCloser
	Resp      *http.Response
}

type Client struct {
	method      string
	urlPoint    string
	urlSite     string
	jsonData    []byte
	dataForm    *url.Values
	httpHeaders http.Header
	httpRequest *http.Request
	httpClient  http.Client
	errorArray  []error
	Cookie      *http.Cookie
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
	SetProxy(proxy *Proxy) *Client
	Method(method string) *Client
	GetMethod() *Client
	PostMethod() *Client
	PutMethod() *Client
	UrlPoint(urlPoint string) *Client
	GetUrl() string
	UrlSite(urlSite string) *Client

	Request() *Client
	Send() HttpResultInterface
	NewRequest() HttpResultInterface
	NewUpdateFile(readFile []byte) HttpResultInterface
}
