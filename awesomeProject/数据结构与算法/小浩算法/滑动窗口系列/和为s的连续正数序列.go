package main

import (
	"fmt"
)

//输入一个正整数 target ，输出所有和为 target 的连续正整数序列（至少含有两个数）。
//序列内的数字由小到大排列，不同序列按照首个数字从小到大排列。
//输入：target = 9
//输出：[[2,3,4],[4,5]]

func main() {
	target := 150
	fmt.Println(findContinuousSequence1(target))
	fmt.Println(findContinuousSequence2(target))
}

func findContinuousSequence1(target int) [][]int {
	nums := make([]int, 0)
	res := make([][]int, 0)
	sum := 1
	// 造正整数序列
	for i := 1; i < target/2+2; i++ {
		nums = append(nums, i)
	}

	for left, right := 1, 2; right < len(nums)+2 && left < right; {
		if sum < target {
			sum += right
			right += 1
		} else if sum > target {
			sum -= left
			left += 1
		} else {
			res = append(res, nums[left-1:right-1])
			sum -= left
			left += 1

		}
	}
	return res
}

func findContinuousSequence2(target int) [][]int {
	result := make([][]int, 0)
	i := 1
	j := 1
	win := 0
	arr := make([]int, target)
	for i := range arr {
		arr[i] = i + 1
	}
	for i <= target/2 {
		if win < target {
			win += j
			j++
		} else if win > target {
			win -= i
			i++
		} else {
			result = append(result, arr[i-1:j-1])
			win -= i
			i++
		}
	}
	return result
}
