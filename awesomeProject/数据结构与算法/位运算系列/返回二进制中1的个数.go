package main

import "fmt"

// 编写一个函数，输入是一个无符号整数，返回其二进制表达式中数字位数为 ‘1’ 的个数（也被称为汉明重量）

func main() {
	fmt.Println(hanMingWeight1(15))
	fmt.Println(hanMingWeight2(15))
}

// 每一位都用1去&
func hanMingWeight1(num int) int {
	res := 0
	compare := 1
	// 64位系统int为64位
	for i := 0; i < 64; i++ {
		if compare&num != 0 {
			res += 1
		}
		compare = compare << 1
	}
	return res
}

// 用 n &(n-1) 去将所有1变0
func hanMingWeight2(num int) int {
	res := 0
	for num > 0 {
		num &= num - 1
		res += 1
	}
	return res
}
