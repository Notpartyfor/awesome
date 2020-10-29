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

	b := []rune(str)
	l := len(b)
	res := make([]string, l)
	period := 2*n - 2
	for i := 0; i < l; i++ {
		mod := i % period

		if mod < n {
			res[mod] += string(b[i])
		} else {
			res[period-mod] += string(b[i])
		}
	}
	return strings.Join(res, "")
}
