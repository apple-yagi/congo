package main

import (
	"fmt"
	"net/http"

	"github.com/apple-yagi/congo/handler"
)

func makeRequest(url string, response *http.Response) handler.GenericFunction {
	return func() error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		fmt.Println(resp)
		response = resp
		return nil
	}
}

func main() {
	var resp1 http.Response
	var resp2 http.Response
	var resp3 http.Response
	funcs := []handler.GenericFunction{
		makeRequest("https://example.com", &resp1),
		makeRequest("https://example2.com", &resp2),
		makeRequest("https://example3.com/", &resp3),
	}
	fmt.Println(handler.RunAsyncAllowErrors(funcs...))
}
