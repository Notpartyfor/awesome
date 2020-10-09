package main

import "fmt"

const IntMax = int(^uint(0) >> 1)
const IntMin = ^IntMax

func main() {
	nums := []int{1, -3, 4, -1, 2, -4, 6, 8, -9, 4, 1}
	k := 4
	fmt.Println(maxSlidingWindow1(nums, k))
	fmt.Println(maxSlidingWindow2(nums, k))
}

// 暴力解法
func maxSlidingWindow1(nums []int, k int) []int {
	res := make([]int, 0)
	l := len(nums)
	// 一共会有 l - k + 1 个窗口
	// 对每个窗口进行遍历
	for i := 0; i < l-k+1; i++ {
		temp := IntMin
		for j := i; j < i+k; j++ {
			temp = max(temp, nums[j])
		}
		res = append(res, temp)
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 双端队列
func maxSlidingWindow2(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	//用切片模拟一个双端队列
	queue := make([]int, 0)
	result := make([]int, 0)
	for i := range nums {
		for i > 0 && (len(queue) > 0) && nums[i] > queue[len(queue)-1] {
			//将比当前元素小的元素祭天
			queue = queue[:len(queue)-1]
		}
		//将当前元素放入queue中
		queue = append(queue, nums[i])
		if i >= k && nums[i-k] == queue[0] {
			//维护队列，保证其头元素为当前窗口最大值
			queue = queue[1:]
		}
		if i >= k-1 {
			//放入结果数组
			result = append(result, queue[0])
		}
	}
	return result
}
