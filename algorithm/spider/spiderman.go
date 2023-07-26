package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetUrl() {
	response, err := http.Get("https://www.mhqwe.xyz/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	document.Find("div.col-xs-12").Each(func(i int, selection *goquery.Selection) {
		//name := selection.Find("img[src]").First()//
		selection.Find("div.thumb-overlay-albums").Each(func(i int, selection *goquery.Selection) {
			url, _ := selection.Find("a").Attr("onclick")
			fmt.Println(url)
			img, _ := selection.Find("a>img").Attr("data-original")
			fmt.Println(img)
		})

	})
}
