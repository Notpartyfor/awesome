package main

import (
	"awesomeProject/common"
	"fmt"
)

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(nums))
}

func maxArea(nums []int) int {
	// 定义双指针
	l, r := 0, len(nums)-1
	res := 0
	for l < r {
		res = common.Max(res, common.Min(nums[l], nums[r])*(r-l))
		if nums[l] <= nums[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
