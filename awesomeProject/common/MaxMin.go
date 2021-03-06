package common

import "strings"

const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func MaxInSlice(nums []int) int {
	max := nums[0]
	for _, v := range nums {
		max = Max(max, v)
	}
	return max
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func MinInSlice(nums []int) int {
	min := nums[0]
	for _, v := range nums {
		min = Min(min, v)
	}
	return min
}

func Abs(num int) int {
	if num < 0 {
		return -1 * num
	}
	return num
}

func StrInStr(source, target string) bool {
	if strings.Index(source, target) == -1 {
		return false
	}
	return true
}

func SumIntSlice(nums []int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
