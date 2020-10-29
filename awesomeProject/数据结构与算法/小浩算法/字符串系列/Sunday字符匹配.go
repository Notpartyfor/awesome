package main

import (
	"fmt"
	"strings"
)

func main() {
	//str := "little"
	str := "here was a little hao"
	pat := "little"
	fmt.Println(Sunday(str, pat))
}

func Sunday(origin, aim string) int {

	if len(origin) < 1 || len(aim) < 1 {
		return 0
	}

	if len(origin) < len(aim) {
		return -1
	}

	originIndex := 0
	aimIndex := 0

	for aimIndex < len(aim) {
		if originIndex > len(origin)-1 {
			return -1
		}
		if origin[originIndex] == aim[aimIndex] {
			originIndex++
			aimIndex++
		} else {
			nextCharIndex := originIndex - aimIndex + len(aim)
			if nextCharIndex < len(origin) {
				step := strings.LastIndex(aim, string(origin[nextCharIndex]))
				if step == -1 {
					originIndex = nextCharIndex + 1
				} else {
					originIndex = nextCharIndex - step
				}
			} else {
				return -1
			}
		}
	}
	return originIndex - aimIndex
}
