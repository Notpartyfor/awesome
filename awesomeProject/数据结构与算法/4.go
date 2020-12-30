package main

import (
	"fmt"
)

func main() {
	nums1 := []int{1, 3}
	nums2 := []int{2, 4, 5, 6, 7, 8}
	fmt.Println(findMedianSortedArrays1(nums1, nums2))
	fmt.Println(findMedianSortedArrays2(nums1, nums2))
}

// 无脑解法，合并在一起然后再找
func findMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	var res float64
	total := len(nums1) + len(nums2)
	merged := make([]int, 0)
	i, j := 0, 0

	for index := 0; index < total; index++ {
		if i == len(nums1) || j == len(nums2) {
			break
		}

		if nums1[i] <= nums2[j] {
			merged = append(merged, nums1[i])
			i++
			continue
		} else {
			merged = append(merged, nums2[j])
			j++
			continue
		}
	}

	if i == len(nums1) {
		merged = append(merged, nums2[j:]...)
	} else {
		merged = append(merged, nums1[i:]...)
	}
	if total%2 == 0 {
		res = (float64(merged[total/2]) + float64(merged[(total/2)-1])) / 2
	} else {
		res = float64(merged[total/2])
	}
	return res
}

// 优化，实际上我们并不需要真的合并在一起，我们先判断出我们需要的数是第几个，然后第几次查询的时候取出来即可
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	var isSingle bool
	var tarNum int
	var result float64
	l1, l2 := len(nums1), len(nums2)
	if (l1+l2)%2 == 1 {
		isSingle = true
	}
	if isSingle {
		tarNum = (l1+l2-1)/2 + 1
	} else {
		tarNum = (l1 + l2) / 2
	}

	var index int
	for {
		index++

		//边界情况 => 在到达tarNum之前某一个空了
		//这样的话，tarNum必然在另一个数组之中，可以直接计算返回
		if l1 == 0 || l2 == 0 {
			var tarLi []int
			if l1 == 0 {
				tarLi = nums2
			} else {
				tarLi = nums1
			}

			if isSingle {
				return float64(tarLi[tarNum-index])
			} else {
				return float64(tarLi[tarNum-index+1]+tarLi[tarNum-index]) / 2
			}
		}

		//正常情况
		//直接把头去掉
		//l1和l2长度也会跟着变化
		if nums1[0] < nums2[0] {
			if index == tarNum {
				result = float64(nums1[0])
			}
			nums1 = nums1[1:]
			l1--
		} else {
			if index == tarNum {
				result = float64(nums2[0])
			}
			nums2 = nums2[1:]
			l2--
		}
		//到达预定情况
		if index == tarNum {
			if isSingle {
				return result
			} else if l1 > 0 && l2 > 0 { // 两个中间数
				if nums1[0] <= nums2[0] {
					return (result + float64(nums1[0])) / 2
				} else {
					return (result + float64(nums2[0])) / 2
				}
			} else if l1 == 0 {
				return (result + float64(nums2[0])) / 2
			} else {
				return (result + float64(nums1[0])) / 2
			}
			break
		}

	}
	return result
}
