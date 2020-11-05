package main

import (
	"fmt"
	"strings"
)

var strs = []string{"abcde", "abcd", "ab"}

//var none = []string{"abcd","efgh"}
//var empty = []string{}

func main() {
	fmt.Println(strs[0])
	fmt.Println(strs[1])
	fmt.Println(strs[2])
	fmt.Println(strings.Index(strs[0], strs[1])) // 如果后者是前者的最长公共前缀，返回0，反之返回-1

	//fmt.Println(longestCommonPrefixMine(strs))
	//fmt.Println(longestCommonPrefixMine(none))
	//fmt.Println(longestCommonPrefixMine(empty))
	//fmt.Println(longestCommonPrefix(strs))
	//fmt.Println(longestCommonPrefix(none))
	//fmt.Println(longestCommonPrefix(empty))
}
func longestCommonPrefixMine(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	temp := strs[0]
	if len(strs) == 1 {
		return strs[0]
	}
	// 用strs[0]作为存储公共前缀
	for i := 1; i < len(strs); i++ {
		// 往后逐个数组对比
		for j := len(strs[0]); j > 0; j-- {
			if strings.HasPrefix(strs[i], strs[0][:j]) {
				// 修改strs[0]
				strs[0] = strs[0][:j]
				break
			}
		}
	}
	// 如果没有变化
	if strs[0] == temp {
		return ""
	}
	return strs[0]
}

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	// 用strs[0]作为存储公共前缀
	prefix := strs[0]
	for _, k := range strs {
		// 如果没包含这个子字符串则会返回-1
		for strings.Index(k, prefix) != 0 {
			if len(prefix) == 0 {
				return ""
			}
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}
