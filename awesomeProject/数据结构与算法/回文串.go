package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	str := "A man, a plan, a canal: Panama"
	str = strings.ToLower(str)
	fmt.Println(str)
	reg := regexp.MustCompile("[^0-9a-z]")
	str = reg.ReplaceAllString(str,"")
	fmt.Println(str)

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


