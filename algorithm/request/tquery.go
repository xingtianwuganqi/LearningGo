package request

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetQuery() {
	url := "https://bbs.hupu.com/love" // 目标网站的URL
	// 发送HTTP GET请求
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// 解析HTML
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//// 提取所需的数据
	//document.Find("h1").Each(func(index int, element *goquery.Selection) {
	//	// 输出标题文本
	//	fmt.Println(element.Text())
	//})

	//document.Find("a").Each(func(index int, element *goquery.Selection) {
	//	// 提取链接和链接文本
	//	link, _ := element.Attr("href")
	//	text := element.Text()
	//
	//	// 输出链接和链接文本
	//	fmt.Println(text, "=>", link)
	//})
	//document.Find(".post-title").Each(func(index int, element *goquery.Selection) {
	//	fmt.Println(element.Text())
	//})
	document.Find("li.bbs-sl-web-post-body").Each(func(index int, element *goquery.Selection) {
		name := element.Find(".post-title").Text()
		// 获取属性值
		value, ok := element.Find("a").Attr("href")
		if ok == false {
			log.Fatal(ok)
		}
		auth := element.Find(".post-auth").Text()
		time := element.Find(".post-time").Text()
		dic := make(map[string]string)
		dic["name"] = name
		dic["url"] = value
		dic["auth"] = auth
		dic["time"] = time

		fmt.Print(dic)
		fmt.Println("\n")
	})
}

// 爬取详情
func detialQuery(url string) {

}
