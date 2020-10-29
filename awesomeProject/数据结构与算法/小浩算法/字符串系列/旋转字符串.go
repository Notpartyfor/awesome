package main

import "fmt"

func main() {
	fmt.Println(rotateString("abcde", "bcdea"))
}

func rotateString(A, B string) bool {
	N := len(A)
	if len(B) != N {
		return false
	}
	if N == 0 {
		return true
	}

	shifts := make([]int, N+1)
	for i, _ := range shifts {
		shifts[i] = 1
	}

	left := -1
	for right := 0; right < N; right++ {
		for left > 0 && B[left] != B[right] {
			left -= shifts[left]
		}
		shifts[right+1] = right - left
		left++
	}

	matchLen := 0
	C := A + A
	for i, _ := range C {
		for matchLen >= 0 && B[matchLen] != C[i] {
			matchLen -= shifts[matchLen]
		}
		matchLen++
		if matchLen == N {
			return true
		}
	}
	return false
}
