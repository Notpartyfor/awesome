package main

import "fmt"

func main() {
	var nums = []int{1, 9, 9}
	fmt.Println(plusOne(nums))
}

func plusOne(digits []int) []int {
	var result []int
	addon := 0
	for i := len(digits) - 1; i >= 0; i-- { // 从最低位开始循环
		// 加上后面给的进位
		digits[i] += addon
		addon = 0
		if i == len(digits)-1 {
			digits[i]++
		}

		// 如果有进位
		if digits[i] == 10 {
			addon = 1
			digits[i] = digits[i] % 10
		}
	}

	// 如果最高位进位的话，证明要扩充数组
	if addon == 1 {
		result = make([]int, 1)
		result[0] = 1
		result = append(result, digits...)
	} else {
		result = digits
	}
	return result
}
