package main

import (
	"fmt"
)

func generate() chan int { // 生成2，3，4，5，6...
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}
func filter(ch chan int, prime int) chan int {
	newch := make(chan int)
	go func() {
		for {
			if i := <-ch; i%prime != 0 {
				newch <- i
			}
		}
	}()
	return newch
}
func sieve() chan int {
	out := make(chan int)
	go func() {
		ch := generate()
		for {
			prime := <-ch          // 拿2，3，4，5，6...
			ch = filter(ch, prime) // 将ch变为已筛选某个数的ch
			out <- prime
		}
	}()
	return out
}

func main() {
	out := sieve()
	for {
		fmt.Println(<-out)
	}
}
