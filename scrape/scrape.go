package scrape

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/saintfish/chardet"
	"golang.org/x/net/html/charset"
)

func ScrapeText(resp *http.Response) {
	defer resp.Body.Close()

	buffer, _ := ioutil.ReadAll(resp.Body)

	detector := chardet.NewTextDetector()
	detectResult, _ := detector.DetectBest(buffer)

	bufferReader := bytes.NewReader(buffer)
	reader, _ := charset.NewReaderLabel(detectResult.Charset, bufferReader)

	document, _ := goquery.NewDocumentFromReader(reader)

	result := document.Find("title").Text()
	fmt.Println(result)
}
