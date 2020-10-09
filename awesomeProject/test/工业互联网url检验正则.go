package main

import (
	"fmt"
	"regexp"
)

const regularUrl = "^(http(s)?:\\/\\/)(www\\.)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+(:\\d+)*(\\/\\w+\\.\\w+)*([\\?&]\\w+=\\w*)*$"

func main() {
	url := "http://baidu.com:333/a.a?a=1&b=3"
	fmt.Println(regularUrl)
	fmt.Println(IsUrlValid(url))
}

func IsUrlValid(url string) bool {
	if m, _ := regexp.MatchString(regularUrl, url); !m {
		return false
	}

	return true
}
