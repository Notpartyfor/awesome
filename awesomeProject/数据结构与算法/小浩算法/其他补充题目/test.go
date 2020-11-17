package main

import "fmt"

func main() {
	secret := "1807"
	guess := "7801"
	fmt.Println(getHint(secret, guess))
}
func getHint(secret string, guess string) string {
	m := make(map[uint8]int)
	for i := range secret {
		m[secret[i]] += 1
	}
	a, b := 0, 0
	n := len(guess)
	for i := 0; i < n; i++ {
		if secret[i] == guess[i] {
			a++
			m[secret[i]] -= 1
			continue
		}

		if m[guess[i]] > 0 {
			b++
			m[guess[i]] -= 1
		}
	}
	return fmt.Sprintf("%d%s%d%s", a, "A", b, "B")
}
