// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.package main
package main

import "fmt"

func generate1(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}
func filter1(in, out chan int, prime int) {
	for {
		i := <-in // Receive value of new variable 'i' from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to channel 'out'.
		}
	}
}
func main() {
	sum := 0
	ch := make(chan int) // Create a new channel.
	go generate1(ch)     // Start generate() as a goroutine.
	for {
		prime := <-ch

		if prime > 2 {
			break
		}
		sum += 1
		fmt.Print(prime, "\n")

		ch1 := make(chan int)
		go filter1(ch, ch1, prime)
		ch = ch1
	}
	fmt.Println(sum)
}
