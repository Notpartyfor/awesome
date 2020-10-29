package main

import "fmt"

// 给定一个整数，编写一个函数来判断它是否是 2 的幂次方。
// 对于2的幂次方，均为  最高位为1 其他位为0
// 对于2的幂次方减一， 均为 最高位为0 其他位为1
// 这是可以记忆的技巧，所以直接 n&(n-1)==0 作为判断

func main() {
	fmt.Println(isPowerOfTwo(16))
}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
