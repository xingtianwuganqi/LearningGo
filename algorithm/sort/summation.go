package sort

import "fmt"

func TwoSum(nums []int, target int) []int {
	var arr []int
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println("i:", i)
		fmt.Println("nums", nums[i])
		for j := i + 1; j < len(nums); j++ {
			fmt.Println("j:", j)
			fmt.Println("nums:", nums[j])
			if nums[i]+nums[j] == target {
				arr = append(arr, i)
				arr = append(arr, j)
				break
			}
		}
	}
	return arr
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

}
