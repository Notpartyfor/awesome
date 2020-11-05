package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "A man, a plan, a canal: Panama"
	// 转小写
	str = strings.ToLower(str)
	fmt.Println(str)
	// 正则匹配replace
	// 不属于0-9a-z的reg
	reg := regexp.MustCompile("[^0-9a-z]")
	// 将其置""
	str = reg.ReplaceAllString(str, "")
	fmt.Println(str)
	// 判断是否为回文串
	l := 0
	r := len(str) - 1
	for l < r {
		if str[l] != str[r] {
			fmt.Println(false)
			break
		}
		l++
		r--
	}
}
