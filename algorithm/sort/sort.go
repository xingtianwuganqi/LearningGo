package sort

// https://blog.csdn.net/fanjufei123456/article/details/123474482

// 冒泡排序
/*
1.比较相邻的元素。如果第一个比第二个大，就交换他们两个。
2.对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对。这步做完后，最后的元素会是最大的数。
3.针对所有的元素重复以上的步骤，除了最后一个。
4.持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。

*/

func BubbleSort(x []int) []int {
	for i := 0; i < len(x)-1; i++ {
		for j := 0; j < len(x)-1-i; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j]
			}
		}
	}
	return x
}

// 选择排序

/*
首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
重复第二步，直到所有元素均排序完毕。
*/

func ChoseSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		var minIndex = i
		for j := i + 1; j < len(arr); j++ {
			if arr[minIndex] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

// 插入排序
//插入排序的原理：始终定义第一个元素为有序的，将元素逐个插入到有序排列之中，其特点是要不断的
/*
将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）
*/

func InsertionSort(arr []int) []int {
	//移动数据，空出一个适当的位置，把待插入的元素放到里面去。
	for i := 0; i < len(arr); i++ {
		//temp 为待排元素 i为其位置 j为已排元素最后一个元素的位置（即取下一个元素，在已经排好序的元素序列中从后向前扫描）
		temp := arr[i]
		//当j < 0 时， i 为第一个元素 该元素认为已经是排好序的 所以不进入while循环
		j := i - 1
		for j >= 0 && arr[j] > temp {
			//如果已经排好序的序列中元素大于新元素，则将该元素往右移动一个位置
			arr[j+1] = arr[j]
			j--
		}
		//跳出while循环时，j的元素小于或等于i的元素(待排元素)。插入新元素 i= j+1
		arr[j+1] = temp
	}
	return arr
}

// 快速排序
/*
快速排序（Quick Sort）是由东尼·霍尔所发展的一种排序算法。在平均状况下，排序 n 个项目要Ο(n log n)次比较。
在最坏状况下则需要Ο(n2)次比较，但这种状况并不常见。事实上，快速排序通常明显比其他Ο(n log n) 算法更快，
因为它的内部循环（inner loop）可以在大部分的架构上很有效率地被实现出来。
快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）。

从数列中挑出一个元素，称为 “基准”（pivot），
重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作。
递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。
递归的最底部情形，是数列的大小是零或一，也就是永远都已经被排序好了。虽然一直递归下去，
但是这个算法总会退出，因为在每次的迭代（iteration）中，它至少会把一个元素摆到它最后的位置去。
*/
func QuickSort(arr []int, left, right int) []int {
	if left > right {
		return []int{}
	}
	i := left
	j := right
	// 记录基准书pivoty
	key := arr[i]
	for i < j {
		//首先从右边j开始查找(从最右边往左找)比基准数(key)小的值<---
		for i < j && key <= arr[j] {
			j--
		}
		//如果从右边j开始查找的值[array[j] integerValue]比基准数小，则将查找的小值调换到i的位置
		if i < j {
			arr[i] = arr[j]
		}
		//从i的右边往右查找到一个比基准数小的值时，就从i开始往后找比基准数大的值 --->
		for i < j && arr[i] <= key {
			i++
		}
		//如果从i的右边往右查找的值[array[i] integerValue]比基准数大，则将查找的大值调换到j的位置
		if i < j {
			arr[j] = arr[i]
		}
	}
	//将基准数放到正确的位置，----改变的是基准值的位置(数组下标)---
	arr[i] = key
	//递归排序
	//将i左边的数重新排序
	QuickSort(arr, left, i-1)
	//将i右边的数重新排序
	QuickSort(arr, i+1, right)
	return arr
}

