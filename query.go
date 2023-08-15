package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"reflect"
	"strings"
)

func (c *Client) getJsonByte(v interface{}) []byte {
	jsonData, err := json.Marshal(v)
	if err != nil {
		c.errorArray = append(c.errorArray, err)
		return nil
	}
	return jsonData
}
func (c *Client) JsonQuery(data interface{}) *Client {
	c.Header("Content-Type", ContentTypeJson)
	switch dataAny := data.(type) {
	case url.Values:
		result := make(map[string]string)
		for key, val := range dataAny {
			if len(val) > 0 {
				result[key] = val[0]
			}
		}
		c.jsonData = c.getJsonByte(result)
	case map[string]interface{}, []map[string]string, []map[string]int:
		c.jsonData = c.getJsonByte(dataAny)
	case string:
		if dataAny[:1] == "{" && dataAny[len(dataAny)-1:] == "}" {
			c.jsonData = []byte(dataAny)
		}
	default:
		if reflect.ValueOf(data).Kind() == reflect.Struct {
			c.jsonData = c.getJsonByte(data)
		}
	}
	return c
}

func (c *Client) FormQuery(data interface{}) *Client {
	c.Header("Content-Type", ContentTypeForm)
	switch dataAny := data.(type) {
	case url.Values:
		for k, v := range dataAny {
			c.dataForm.Set(k, v[0])
		}
	case map[string]interface{}:
		for k, v := range dataAny {
			c.dataForm.Set(k, fmt.Sprintf("%v", v))
		}
	case map[string]string:
		for k, v := range dataAny {
			c.dataForm.Set(k, v)
		}
	case string:
		dataForm, err := url.ParseQuery(dataAny)
		if err != nil {
			c.errorArray = append(c.errorArray, err)
		} else {
			for k, v := range dataForm {
				c.dataForm.Set(k, v[0])
			}
		}
	}
	return c
}

func (c *Client) QueryResult() io.Reader {
	if c.jsonData != nil {
		return bytes.NewReader(c.jsonData)
	} else if c.dataForm != nil {
		return bytes.NewReader([]byte(c.dataForm.Encode()))
	} else {
		return nil
	}
}
func (c *Client) Query(data interface{}, queryType string) *Client {
	if queryType == ContentTypeFormString {
		c.FormQuery(data)
	} else if queryType == ContentTypeJsonString {
		c.JsonQuery(data)
	} else {
		c.errorArray = append(c.errorArray, fmt.Errorf("query error: form and json is false"))
	}
	return c
}

func (c *Client) QueryFunc(f func(c *Client) (interface{}, string)) *Client {
	data, queryType := f(c)
	if data != nil {
		if !strings.Contains(queryType, strings.Join([]string{ContentTypeJsonString, ContentTypeFormString}, "|")) {
			c.errorArray = append(c.errorArray, fmt.Errorf("queryType error: %s is not support", queryType))
		} else {
			c.Query(data, queryType)
		}
	} else {
		c.errorArray = append(c.errorArray, fmt.Errorf("error: %s", "QueryFunc return nil"))
	}
	return c
}
