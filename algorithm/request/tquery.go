package request

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetQuery() {
	resp, err := http.Get("https://bbs.hupu.com/all-gambia")
	if err != nil {
		log.Fatal(err)
	}
	if resp.StatusCode != 200 {
		panic(resp)
	}
	defer resp.Body.Close()

	doc, docErr := goquery.NewDocumentFromReader(resp.Body)
	if docErr != nil {
		log.Fatal(docErr)
	}
	fmt.Println(doc)
}
