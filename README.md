# API Documentation

## Package client

The `requests` package provides a set of interfaces and structs for making HTTP requests and handling responses.

### Structs

#### Proxy

- `Ip` (string): The IP address of the proxy server.
- `Port` (string): The port number of the proxy server.
- `UserName` (string): The username for authentication with the proxy server.
- `Password` (string): The password for authentication with the proxy server.

#### Response

- `BodyBytes` ([]byte): The response body as a byte slice.
- `Body` (io.ReadCloser): An `io.ReadCloser` that allows reading the response body.
- `Resp` (*http.Response): The underlying `http.Response` object.

#### Client

- `method` (string): The HTTP request method (e.g., GET, POST, PUT).
- `urlPoint` (string): The endpoint path of the URL.
- `urlSite` (string): The base URL of the site.
- `DataForm` (*url.Values): The form data for the request.
- `httpHeaders` (http.Header): The HTTP headers for the request.
- `httpRequest` (*http.Request): The underlying `http.Request` object.
- `httpClient` (http.Client): The HTTP client used to send the request.
- `errorArray` ([]error): An array to store any encountered errors.
- `Cookie` (*http.Cookie): The HTTP cookie for the request.

### Interfaces

#### HttpResultInterface

- `SetDecodeFunc(func(f *Response) error) *Response`: Sets the decoding function for the response.
- `Bytes() []byte`: Returns the response body as a byte slice.
- `String() string`: Returns the response body as a string.
- `Json() gjson.Result`: Parses the response body as JSON and returns the result.
- `Dict() map[string]interface{}`: Parses the response body as JSON and returns a dictionary.
- `Decode(v any) error`: Decodes the response body into the provided value.
- `DecodePrintError(v any)`: Decodes the response body into the provided value and prints any encountered errors.
- `GetCookie() []*http.Cookie`: Returns the cookies received in the response.
- `GetHeader() http.Header`: Returns the headers received in the response.
- `GetStatusCode() int`: Returns the status code of the response.
- `Document() *goquery.Document`: Parses the response body as HTML and returns a `goquery.Document`.
- `GetStatus() string`: Returns the status of the response.

#### HttpClientInterface

- `Query(q map[string]interface{}) *Client`: Sets the query parameters for the request.
- `QueryFunc(f func(c *Client)) *Client`: Sets the query parameters for the request using a function.
- `Headers(h map[string]interface{}) *Client`: Sets the HTTP headers for the request.
- `HeadersFunc(f func(c *Client)) *Client`: Sets the HTTP headers for the request using a function.
- `SetCookie(cookie map[string]string) *Client`: Sets the HTTP cookie for the request.
- `SetProxy(proxy Proxy) *Client`: Sets the proxy for the request.
- `Method(method string) *Client`: Sets the HTTP request method.
- `GetMethod() *Client`: Sets the HTTP request method to GET.
- `PostMethod() *Client`: Sets the HTTP request method to POST.
- `PutMethod() *Client`: Sets the HTTP request method to PUT.
- `UrlPoint(urlPoint string) *Client`: Sets the endpoint path of the URL.
- `GetUrl() string`: Returns the complete URL for the request.
- `UrlSite(urlSite string) *Client`: Sets the base URL of the site.
- `Request() *Client`: Returns the underlying `Client` object.
- `Send() HttpResultInterface`: Sends the HTTP request and returns the response as an `HttpResultInterface`.
- `NewRequest() HttpResultInterface`: Creates a new HTTP request and returns the response as an `HttpResultInterface`.
