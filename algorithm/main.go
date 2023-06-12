package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	fmt.Println("哈哈哈")
	arr := []int{10, 1, 9, 5, 7, 3, 4}
	bubbleArr := sort.BubbleSort(arr)
	fmt.Println(bubbleArr)
	choseArr := sort.ChoseSort(arr)
	fmt.Println(choseArr)
	insertArr := sort.InsertionSort(arr)
	fmt.Println(insertArr)
}
