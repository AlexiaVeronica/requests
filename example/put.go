package main

import (
	"fmt"
	"github.com/VeronicaAshford/requests"
)

type Data struct {
	Key  string `json:"key"`
	Page int    `json:"page"`
	Size int    `json:"size"`
}

func main() {
	params := map[string]any{"key": "", "page": 1, "size": 5}
	headers := map[string]interface{}{"Content-Type": requests.ContentTypeForm}
	response := requests.Put("http://localhost:8080/api/v1/Demo/GetDemo", params, headers)
	fmt.Println(response.String())

	params2 := Data{Key: "", Page: 1, Size: 5}
	headers2 := map[string]interface{}{"Content-Type": requests.ContentTypeJson}
	response2 := requests.Put("http://localhost:8080/api/v1/Demo/GetDemo", params2, headers2)
	fmt.Println(response2.String())

}
