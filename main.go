package main

import (
	"fmt"
	"net/http"

	"github.com/apple-yagi/congo/handler"
	"github.com/apple-yagi/congo/scrape"
)

func main() {
	var resp1 http.Response
	var resp2 http.Response
	var resp3 http.Response
	funcs := []handler.GenericFunction{
		makeRequest("https://example.com", &resp1),
		makeRequest("https://example2.com", &resp2),
		makeRequest("https://example3.com/", &resp3),
	}

	errs := (handler.RunAsyncAllowErrors(funcs...))
	for i := range errs {
		if errs[i] != nil {
			fmt.Println("An error occurred in resp", i, ": ", errs[i])
		}
	}
}

func makeRequest(url string, response *http.Response) handler.GenericFunction {
	return func() error {
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		response = resp

		// Scraping response
		scrape.ScrapeText(response)

		return nil
	}
}
