package main

import (
	"fmt"
	"github.com/VeronicaAshford/requests"
)

func main() {
	params := map[string]any{"key": "", "page": 1, "size": 5}
	headers := map[string]interface{}{"Content-Type": requests.ContentTypeForm}
	response := requests.Post("http://localhost:8080/api/v1/Demo/GetDemo", params, headers)
	fmt.Println(response.String())

	params2 := Data{Key: "", Page: 1, Size: 5}
	headers2 := map[string]interface{}{"Content-Type": requests.ContentTypeJson}
	response2 := requests.Post("http://localhost:8080/api/v1/Demo/GetDemo", params2, headers2)
	fmt.Println(response2.String())

}
