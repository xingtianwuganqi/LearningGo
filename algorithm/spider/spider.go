package spider

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetHttp() {
	resp, error := http.Get("https://www.wai77.com/tag/%E7%A7%80%E4%BA%BA%E7%BD%91/")
	if error != nil {
		log.Fatal(error)
		return
	}
	fmt.Println("++++++")
	defer resp.Body.Close()
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("=====")
	document.Find("#recent-content").Each(func(i int, selection *goquery.Selection) {
		selection.Find("div[id]").Each(func(i int, selection *goquery.Selection) {
			fmt.Println(selection)
			idInfo, ok := selection.Attr("id")
			if ok {
				fmt.Println(idInfo)
			}
			img, ok := selection.Find("a.thumbnail-link").First().Attr("href")
			if ok {
				fmt.Println(img)
			}
		})
		//src, ok := selection.Find("img").Attr("src")
		//if ok {
		//	fmt.Println(src)
		//}
	})
}
