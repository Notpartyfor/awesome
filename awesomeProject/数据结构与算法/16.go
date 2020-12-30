package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3, 0, 1, 2}
	target := 1
	fmt.Println(threeSumClosest(nums, target))
}
func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	// 先判断target是否处于三数和边界外
	// 1. 如果target比最小的组合还小或者比最大的组合还大，直接返回
	if target <= nums[0]+nums[1]+nums[2] {
		return nums[0] + nums[1] + nums[2]
	} else if target >= nums[len(nums)-1]+nums[len(nums)-2]+nums[len(nums)-3] {
		return nums[len(nums)-1] + nums[len(nums)-2] + nums[len(nums)-3]
	}

	// 此时target就在最小和最大的组合之间了
	// 以前三个数 - target 作为初始值
	res := nums[0] + nums[1] + nums[2] - target

	for i := 0; i < n-1; i++ {
		l, r := i+1, n-1
		if i == 0 || nums[i] != nums[i-1] {
			for l < r {
				temp := nums[i] + nums[l] + nums[r] - target
				// 相等就直接返回
				if temp == 0 {
					return target
				}
				// 如果大于
				if temp > 0 {
					r--
				} else if temp < 0 {
					l++
				}
				// 如果距离更小了
				if abs(temp) < abs(res) {
					res = temp
				}
			}
		}
	}
	return res + target
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
