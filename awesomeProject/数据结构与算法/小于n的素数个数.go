package main

import "fmt"

func main() {
	fmt.Println(countPrimes3(10))
}

// 超时枚举
func countPrimes1(n int) (cnt int) {
	isPrime := func(x int) bool {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false
			}
		}
		return true
	}
	for i := 2; i < n; i++ {
		if isPrime(i) {
			fmt.Println(i)
			cnt++
		}
	}
	return
}

// 埃氏筛
func countPrimes2(n int) (cnt int) {
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		if isPrime[i] {
			cnt++
			// 将其各个倍数的置为合数
			// 显然，这会重复标记
			for j := 2 * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}
	return
}

// 线性筛
func countPrimes3(n int) int {
	primes := make([]int, 0)
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}
	for i := 2; i < n; i++ {
		// 如果是素数，加进去prime数组
		if isPrime[i] {
			primes = append(primes, i)
		}

		for _, p := range primes {
			// 每个素数乘以i那绝对是合素
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			// 让每个合数只被标记一次
			// 如果i可以被prime[j]整除，那么对于合数 y = i * prime[j+1] 来说
			// (x / prime[j])
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}
