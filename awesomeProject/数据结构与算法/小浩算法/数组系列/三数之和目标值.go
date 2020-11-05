package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}, 3))
}

func threeSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < n; i++ {
		l := i + 1
		r := n - 1
		// nums[i]和nums[i-1]相等的时候跳过对比
		if i == 0 || nums[i] != nums[i-1] {
			// 第一个数不变，处理第二第三个数
			for l < r {
				// 相等的话就加到res返回数组里面
				if nums[i]+nums[l]+nums[r] == target {
					res = append(res, []int{nums[i], nums[l], nums[r]})
					// 同时跳过相同的
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					// 推进
					l++
					r--
				} else if nums[i]+nums[l]+nums[r] > target {
					r--
				} else if nums[i]+nums[l]+nums[r] < target {
					l++
				}
			}
		}
	}
	return res
}
