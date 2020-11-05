package main

import (
	"fmt"
	"sort"
)

func main() {
	var nums1 = []int{1, 3, 2, 2, 9}
	var nums2 = []int{2, 3, 9, 6, 5}
	fmt.Println(intersect(nums1, nums2))
	sort.Ints(nums1)
	sort.Ints(nums2)
	fmt.Println(intersectSoft(nums1, nums2))
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] += 1
	}
	k := 0
	for _, v := range nums2 {
		if m[v] > 0 {
			m[v] -= 1
			// 顺便将nums2 承载交集，作为结果数组返回
			nums2[k] = v
			k++
		}
	}
	return nums2[:k]
}

// 如果数组已经排序好，可以优化
func intersectSoft(nums1 []int, nums2 []int) []int {
	i, j, k := 0, 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			j++
		} else if nums2[j] > nums1[i] {
			i++
		} else {
			nums1[k] = nums1[i]
			i++
			j++
			k++
		}
	}
	return nums1[:k]
}
