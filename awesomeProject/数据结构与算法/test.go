package main

import "fmt"

func main() {
	ch := make(chan int)
	go gen(ch)
	for m := range ch {
		fmt.Println(m)
	}
}
func gen(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i

	}
}
