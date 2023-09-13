package main

import (
	"fmt"
	"github.com/VeronicaAshford/requests"
)

func main() {
	params := map[string]any{"key": "", "page": 1, "size": 5}
	headers := map[string]any{"Content-Type": requests.ContentTypeForm}
	response := requests.Get("http://localhost:8080/api/v1/Demo/GetDemo", params, headers)
	fmt.Println(response.String())

}
