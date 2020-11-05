package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "LEETCODEISHIRING"
	fmt.Println(convert(str, 3))
}
func convert(str string, n int) string {
	if n == 1 {
		return str
	}
	// 用rune结构使得汉字也可以冲
	b := []rune(str)
	l := len(b)
	res := make([]string, l)
	// 每2n-2个周期会回到Z的同一个位置
	period := 2*n - 2
	for i := 0; i < l; i++ {
		// 这个是确认在一个周期的阶段
		mod := i % period
		// 如果小于n，证明还未反向，顺着加
		if mod < n {
			res[mod] += string(b[i])
		} else { // 反向后则需要相减一下
			res[period-mod] += string(b[i])
		}
	}
	return strings.Join(res, "")
}
