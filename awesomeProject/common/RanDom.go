package common

func RandomIntSlice(n int) []int {
	res := make([]int, 0)
	for i := range random(n) {
		res = append(res, i)
	}
	return res
}
func random(n int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < n; i++ {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}()
	return c
}
