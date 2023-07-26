package spider

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sndnvaps/webp2jpg"
	"image"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
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
			idInfo, ok := selection.Attr("id")
			if ok {
				fmt.Println(idInfo)
			}
			img, ok := selection.Find("a.thumbnail-link").First().Attr("href")
			if ok {
				fmt.Println(img)
				SpiderDetail(img)
			}
		})
	})
}

func SpiderDetail(url string) {
	resp, error := http.Get(url)
	if error != nil {
		log.Fatal(error)
		return
	}
	defer resp.Body.Close()
	fmt.Println("+++++", resp.Body)
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("-----")
	//.spotlight text-center
	document.Find("div.spotlight.text-center").Each(func(i int, selection *goquery.Selection) {
		fmt.Println(selection)
		img, ok := selection.Attr("data-src")
		if ok {
			imgUrl := "https://www.wai77.com" + img
			fmt.Println(imgUrl)
			DownLoadImg(imgUrl, "xrw")
		}
	})
}

func DownLoadImg(imgUrl, path string) {
	// 创建本地文件夹

	names := strings.Split(imgUrl, "/")
	var name string = ""
	var newName string = ""
	var newPath string = ""
	if len(names) > 0 {
		name = names[len(names)-1]
		newPath = path + "/" + names[len(names)-2]
		newName = strings.Replace(name, ".webp", ".jpg", 1)
		newName = newPath + "/" + newName
	}
	fmt.Println(name)
	response, error := http.Get(imgUrl)
	if error != nil {
		log.Fatal(error, "1")
		return
	}
	defer response.Body.Close()
	// 看是否能打开文件夹，
	_, error = os.Open(newPath)
	if error != nil {
		os.Mkdir(newPath, 0777)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err, "2")
		return
	}
	//err = ioutil.WriteFile(name, content, 0666)
	img, _, err := image.Decode(bytes.NewReader(content))
	fmt.Println(newName)
	webp2jpg.Encode(img, newName, "jpg")
	if err != nil {
		log.Fatal(err)
		return
	}
}
