package requests

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"net/http"
)

type Request struct {
	url         string
	httpRequest *http.Request
	httpClient  http.Client
}
type HttpRequestInterface interface {
	Send() HttpResultInterface
	NewUpdateFile(readFile []byte) HttpResultInterface
	Stream() chan []byte
}

func (c *Request) Send() HttpResultInterface {
	if c.httpRequest == nil {
		return NewResponse(nil)
	}
	respClient, err := c.httpClient.Do(c.httpRequest)
	if err != nil {
		log.Printf("httpClient.Do error: %v", err)
	}
	return NewResponse(respClient)

}

func (c *Request) Stream() chan []byte {
	streamChan := make(chan []byte)
	resp, ok := c.httpClient.Do(c.httpRequest)
	if ok != nil {
		log.Printf("httpClient.Do error: %v", ok)
	} else {
		go func() {
			reader := bufio.NewReader(resp.Body)
			for {
				line, err := reader.ReadBytes('\n')
				if err != nil {
					if err == io.EOF {
						break
					}
					log.Printf("reader.ReadBytes error: %v", err)
				}
				streamChan <- line
			}
		}()
	}
	return streamChan
}
func (c *Request) NewUpdateFile(readFile []byte) HttpResultInterface {
	res, err := http.Post(c.url, "multipart/form-data", bytes.NewReader(readFile))
	if err != nil {
		log.Printf("http.Post error: %v", err)
	} else {
		return &Response{Body: res.Body, Resp: res}
	}
	return HttpResultInterface(nil)
}
