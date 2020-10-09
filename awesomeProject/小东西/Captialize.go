package main

import "fmt"

func main() {
	vName := "users"
	vType := "[]User"
	result := "[]" + Capitalize(vName)

	fmt.Println(result[0:len(result)-1])
	fmt.Println(result==vType+"s" || result==vType+"es")
}

func Capitalize(str string) string{
	vv := []rune(str)
	if vv[0] >= 97 && vv[0] <= 122 {
		vv[0] -= 32
	}
	return string(vv)
}
