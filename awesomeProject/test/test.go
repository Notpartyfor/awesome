package main

import (
	"fmt"
	"time"
)

func main() {
	var timeS int64 = 1603942026
	t := time.Unix(timeS, 0).Format("2006-01-02 15:04:05")
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(t)
}
