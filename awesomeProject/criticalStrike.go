package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	t1 := time.Now().Second()
	ch := make(chan int, 10000)
	for i := 0; i < 16; i++ {
		go Try(ch)
	}
	total, yes := 0, 0
	for {
		t := <-ch
		total += t
		yes += 1
		if total == 1e7 {
			break
		}
	}

	fmt.Println(total, yes)
	t2 := time.Now().Second()
	fmt.Println(t2 - t1)
}

func Try(out chan int) {
	for {
		base := 39
		total := 0
		for {
			total += 1
			if !isCriticalStrike(base) {
				if base == 79 {
					continue
				}
				base += 8
			} else {
				break
			}

		}
		out <- total
	}
}

func isCriticalStrike(base int) bool {
	if base >= rand.Intn(100) {
		return true
	}
	return false
}
