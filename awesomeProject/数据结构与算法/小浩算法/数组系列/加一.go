package main

import "fmt"

func main() {
	var nums = []int{2, 3, 6, 5, 4, 1, 9, 9}
	fmt.Println(plusOne(nums))
}

func plusOne(digits []int) []int {
	var result []int
	// 进位
	addon := 0
	for i := len(digits) - 1; i >= 0; i-- { // 从最低位开始循环
		// 加上后面给的进位
		digits[i] += addon
		// 进位置零
		addon = 0
		// 初始加一
		if i == len(digits)-1 {
			digits[i]++
		}

		// 如果有进位
		if digits[i] == 10 {
			addon = 1
			digits[i] = digits[i] % 10
		}
	}
	// 上述循环一直循环到i=0后，也就是最高位
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
