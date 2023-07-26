package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	fmt.Println("哈哈哈")
	arr := []int{10, 1, 9, 5, 7, 3, 4}
	//bubbleArr := sort.BubbleSort(arr)
	//fmt.Println(bubbleArr)
	//choseArr := sort.ChoseSort(arr)
	//fmt.Println(choseArr)

	nums := sort.TwoSum(arr, 11)
	fmt.Println("nums:")
	fmt.Println(nums)
	//insertArr := sort.InsertionSort(arr)
	//fmt.Println(insertArr)
	//
	//quickArr := sort.QuickSort(arr, 0, len(arr)-1)
	//fmt.Println(quickArr)

	//request.GetQuery()
	//request.DetialQuery("60585128.html")
	//request.DownLoadImg("https://i5.hoopchina.com.cn/hupuapp/bbs/917/99291917/1686564522_5091228IMG_5692.png?x-oss-process=image/resize,w_800/format,webp", "./imgs")
}
