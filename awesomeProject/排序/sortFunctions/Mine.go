package sortFunctions

import (
	"awesomeProject/common"
	"math"
)

// 自己做一遍加深理解呗
// todo 冒泡排序
func BubbleSortMine(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// todo 基数排序
func RadixSortMine(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	maxDigit := 1
	max := common.MaxInSlice(nums)
	for int(math.Pow10(maxDigit)) < max {
		maxDigit += 1
	}

	for Digit := 0; Digit < maxDigit; Digit++ {
		temp := make([][]int, 10)
		for i := range temp {
			temp[i] = make([]int, 0)
		}

		for _, v := range nums {
			t := (v / int(math.Pow10(Digit))) % 10
			temp[t] = append(temp[t], v)
		}

		resultPerRound := make([]int, 0)
		for _, bucket := range temp {
			for _, v := range bucket {
				resultPerRound = append(resultPerRound, v)
			}
		}

		nums = resultPerRound
	}
	return nums
}

// todo 选择排序
func SelectionSortMine(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}
	for i := 0; i < n; i++ {
		minIndex := i
		for j := i + 1; j < n-i-1; j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
	return nums
}

// todo 插入排序
func InsertionSortMine(nums []int) []int {
	n := len(nums)
	if n < 2 {
		return nums
	}

	for i := range nums {
		preIndex := i - 1
		current := nums[i]

		if preIndex >= 0 && nums[preIndex] > current {
			nums[preIndex+1] = nums[preIndex]
			preIndex -= 1
		}

		nums[preIndex+1] = current
	}
	return nums
}
