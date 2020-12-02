package main

import (
	"awesomeProject/common"
	"fmt"
)

func main() {
	str := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring1(str))
	fmt.Println(lengthOfLongestSubstring2(str))

}

// 双指针
func lengthOfLongestSubstring1(str string) int {
	if len(str) == 0 {
		return 0
	}
	l := len(str)
	res := 0
	left := 0
	right := 1
	temp := string(str[left])
	for left < l && right < l {
		target := string(str[right])
		if !common.StrInStr(temp, target) {
			right += 1
			temp = temp + target
			res = common.Max(res, right-left)
		} else {
			temp = temp[1:]
			left += 1
		}

	}
	return res
}

// 用map进行优化
func lengthOfLongestSubstring2(str string) int {
	if len(str) == 0 {
		return 0
	}
	l := len(str)
	result := 0
	m := make(map[string]int)

	// 要把头也塞进去，这里right要赋0的
	for left, right := 0, 0; right < l; right++ {
		// 如果有，就要改left了
		if _, ok := m[string(str[right])]; ok {
			left = common.Max(m[string(str[right])], left)
		}
		result = common.Max(result, right-left+1)
		m[string(str[right])] = right + 1
	}
	return result
}
