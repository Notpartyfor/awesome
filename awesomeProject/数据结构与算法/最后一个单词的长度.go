package main

import "fmt"

func main() {
	fmt.Println(lengthOfLastWord("Hello World"))
}

func lengthOfLastWord(s string) int {
	if len(s) < 1 {
		return 0
	}
	count := 0

	for i := len(s) - 1 ; i > 0 ; i -- {
		if s[i] == ' ' {
			if count == 0 {
				continue
			}
			break
		}
		count ++
	}
	return count
}
