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
