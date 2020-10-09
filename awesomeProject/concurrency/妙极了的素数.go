package main

import (
	"fmt"
)
func generate() chan int {
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
			prime := <-ch
			ch = filter(ch, prime)
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