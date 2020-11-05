package main

import "fmt"

func main() {
	var nums = []int{2, 3, 6, 5, 4, 1, 9, 9}
	fmt.Println(nums)
	rotate(nums, 4)
	fmt.Println(nums)
}

// 切片引用，所以就当成指针了？
func rotate(nums []int, k int) {
	reverse(nums)
	reverse(nums[:k%len(nums)])
	reverse(nums[k%len(nums):])
}

func reverse(arr []int) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
