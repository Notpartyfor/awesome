package main

import "fmt"

// 给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。
// 你可以返回满足此条件的任何数组作为答案。

func main() {
	nums := []int{7, 1, 3, 4, 5}
	fmt.Println(sortArrayByParity(nums))
}

func sortArrayByParity(nums []int) []int {
	j := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}
	return nums
}
