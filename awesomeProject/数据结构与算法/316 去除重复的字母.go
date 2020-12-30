package main

import (
	"fmt"
)

func removeDuplicateLetters(s string) string {
	// 记录每个字母出现的次数，因为最后需要每种只保留一个
	var times [26]int
	for _, v := range s {
		times[v-'a'] += 1
	}

	stack := make([]rune, 0)
	for _, c := range s { // 值用rune(int32)

		// 丢弃逻辑
		// 栈顶元素的丢弃次数 > 1 (最后需要为1) (每个字母不同)
		// 栈中有元素 (跳过第一个，避免后面条件出错)
		// 比较当前和栈顶的序大小
		for len(stack) > 0 && times[stack[len(stack)-1]-'a'] > 1 && c < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			times[c-'a'] -= 1
		}
		stack = append(stack, c)
	}
	return string(stack)
}

func main() {
	s := "bcabc"
	fmt.Println(removeDuplicateLetters(s))
}
