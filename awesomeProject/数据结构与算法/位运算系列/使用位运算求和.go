package main

import "fmt"

// 求 1 2 ... n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
func main() {
	fmt.Println(sumNums(3))
}

func plus(a *int, b int) bool {
	*a += b
	return true
}
func sumNums(n int) int {
	_ = n > 0 && plus(&n, sumNums(n-1))
	return n
}
