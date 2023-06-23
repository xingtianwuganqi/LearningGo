package request

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetQuery() {
	url := "https://bbs.hupu.com/787" // 目标网站的URL
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

		fmt.Print(dic, "\n")
		DetialQuery(value)
	})
}

// 爬取详情
func DetialQuery(url string) {
	newUrl := "https://bbs.hupu.com/" + url
	response, error := http.Get(newUrl)
	if error != nil {
		log.Fatal(error)
	}
	defer response.Body.Close()
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	document.Find(".index_bbs-post-web-main__D_K6v").Each(func(index int, element *goquery.Selection) {
		content := element.Find(".thread-content-detail").Text()
		// 遍历图片
		var image []string
		element.Find(".slate-image").Each(func(index int, element *goquery.Selection) {
			imgUrl, ok := element.Find("img").Attr("src")
			if ok == false {
				log.Fatal(ok)
			}
			image = append(image, imgUrl)
		})
		// 爬取评论列表
		var detail_content []string
		var detail_imgs []string
		element.Find(".m-c").Each(func(index int, element *goquery.Selection) {
			dText := element.Find(".thread-content-detail").Text()
			if len(dText) > 0 {
				detail_content = append(detail_content, dText)
			}
			dImg, ok := element.Find("img").Attr("src")
			if ok && len(dImg) > 0 {
				detail_imgs = append(detail_imgs, dImg)
			}
		})
		fmt.Println(content)
		fmt.Println(image)
		fmt.Println(detail_content)
		fmt.Println(detail_imgs)

		for i := 0; i < len(image); i++ {
			DownLoadImg(image[i], "./imgs")
		}
	})
}

func DownLoadImg(imgUrl, path string) {
	// 创建本地文件夹

	if strings.Contains(imgUrl, "?") {
		imgUrl = strings.Split(imgUrl, "?")[0]
	}

	names := strings.Split(imgUrl, "/")
	var name string = ""
	if len(names) > 0 {
		name = path + "/" + names[len(names)-1]
	}
	fmt.Println(name)
	response, error := http.Get(imgUrl)
	if error != nil {
		log.Fatal(error, "1")
		return
	}
	defer response.Body.Close()
	// 看是否能打开文件夹，
	_, error = os.Open(path)
	if error != nil {
		os.Mkdir(path, 0777)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err, "2")
		return
	}
	err = ioutil.WriteFile(name, content, 0666)
	if err != nil {
		log.Fatal(err)
		return
	}

}
