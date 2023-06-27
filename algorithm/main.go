package main

import (
	"algorithm/spider"
	"fmt"
)

func main() {
	fmt.Println("哈哈哈")
	//arr := []int{10, 1, 9, 5, 7, 3, 4}
	//bubbleArr := sort.BubbleSort(arr)
	//fmt.Println(bubbleArr)
	//choseArr := sort.ChoseSort(arr)
	//fmt.Println(choseArr)
	//insertArr := sort.InsertionSort(arr)
	//fmt.Println(insertArr)
	//
	//quickArr := sort.QuickSort(arr, 0, len(arr)-1)
	//fmt.Println(quickArr)

	//request.GetQuery()
	//request.DetialQuery("60585128.html")
	//request.DownLoadImg("https://i5.hoopchina.com.cn/hupuapp/bbs/917/99291917/1686564522_5091228IMG_5692.png?x-oss-process=image/resize,w_800/format,webp", "./imgs")

	spider.GetHttp()
	//detialUrl := "https://www.wai77.com/xiuren%e7%a7%80%e4%ba%ba%e7%bd%91-no-5751-%e6%9d%a8%e6%99%a8%e6%99%a8yome-%e9%bb%91%e8%89%b2%e5%90%8a%e5%b8%a6%e4%b8%8a%e8%a1%a3%e6%90%ad%e9%85%8d%e7%ba%a2%e8%89%b2%e4%b8%9d%e8%a2%9c/"
	//spider.SpiderDetail(detialUrl)
}
