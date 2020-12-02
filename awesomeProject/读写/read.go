package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("工时记录-11.txt")
	if err != nil {
		fmt.Println(fmt.Errorf("读文件出错, %s", err))
	}
	defer file.Close()

	iReader := bufio.NewReader(file)
	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			return // error or EOF
		}
		fmt.Printf("The input was: %s", str)
	}
}
