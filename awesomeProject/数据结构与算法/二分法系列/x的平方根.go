package main

import "fmt"

// 计算并返回 x 的平方根，其中 x 是非负整数。由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

// 没有小数的话，查找区间就很好确认了

func main() {
	fmt.Println(mySqrt(25))
}

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	left := 0.0
	right := float64(x)

	for left <= right {
		mid := (left + right) / 2
		if mid*mid == float64(x) {
			return int(mid)
		} else if mid*mid < float64(x) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return int(left)
}
