package main

import (
	"fmt"
)

// 给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

func main() {
	nums := []int{2, 3, 1, 0, 5}
	fmt.Println(missingNumber1(nums))
	fmt.Println(missingNumber2(nums))
}

func missingNumber1(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	n := len(nums)
	return n*(n+1)/2 - sum
}

// 位运算(i跟k抵消
func missingNumber2(nums []int) int {
	result := 0
	for i, k := range nums {
		result ^= k ^ i
	}
	return result ^ len(nums)
}
