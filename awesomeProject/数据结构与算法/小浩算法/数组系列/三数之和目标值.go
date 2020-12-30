package main

import (
	"fmt"
	"sort"
)

func main() {
	//fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}, 3))
	fmt.Println(fourSum([]int{-2, -1, 1, 2, 3, 4, 5, 6}, 3))

}

func threeSum(nums []int, target int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	res := make([][]int, 0)

	for i := 0; i < n-1; i++ {
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

func fourSum(nums []int, target int) (quadruplets [][]int) {
	sort.Ints(nums)
	n := len(nums)
	for i := 0; i < n-3 && nums[i]+nums[i+1]+nums[i+2]+nums[i+3] <= target; i++ {
		if i > 0 && nums[i] == nums[i-1] || nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target {
			continue
		}
		for j := i + 1; j < n-2 && nums[i]+nums[j]+nums[j+1]+nums[j+2] <= target; j++ {
			if j > i+1 && nums[j] == nums[j-1] || nums[i]+nums[j]+nums[n-2]+nums[n-1] < target {
				continue
			}
			for left, right := j+1, n-1; left < right; {
				if sum := nums[i] + nums[j] + nums[left] + nums[right]; sum == target {
					quadruplets = append(quadruplets, []int{nums[i], nums[j], nums[left], nums[right]})
					for left++; left < right && nums[left] == nums[left-1]; left++ {
					}
					for right--; left < right && nums[right] == nums[right+1]; right-- {
					}
				} else if sum < target {
					left++
				} else {
					right--
				}
			}
		}
	}
	return
}
