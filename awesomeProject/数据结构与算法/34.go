package main

import (
	"fmt"
)

// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。
//
//如果数组中不存在目标值 target，返回 [-1, -1]。
//
//进阶：
//
//你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
//
//
//示例 1：
//
//输入：nums = [5,7,7,8,8,10], target = 8
//输出：[3,4]
//示例 2：
//
//输入：nums = [5,7,7,8,8,10], target = 6
//输出：[-1,-1]
//

func main() {
	nums := []int{5, 7, 7, 8, 8, 19}
	target := 19
	fmt.Println(find(nums, target))
}

// 自己实现的二分法
func find(nums []int, target int) []int {
	left, right := 0, len(nums)

	// 特别条件快速判断
	if right == 0 {
		return []int{-1, -1}
	}

	if target < nums[left] || target > nums[right-1] {
		return []int{-1, -1}
	}

	// 二分法查找等于mid的值，找不到就返回
	mid := 0
	for left < right {
		mid = (left + right) / 2
		// 找到相等的mid之后向左右延伸
		if nums[mid] == target {
			i, j := mid, mid
			for i > 0 {
				if nums[i-1] == target {
					i -= 1
				} else {
					break
				}
			}
			for j < len(nums)-1 {
				if nums[j+1] == target {
					j += 1
				} else {
					break
				}
			}
			return []int{i, j}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return []int{-1, -1}
}
