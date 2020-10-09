package main

import "fmt"
func main() {

	nums := [][]int{
		{7,3,1,9},
		{3,4,6,9},
		{6,9,6,6},
		{9,5,8,5},
	}

	fmt.Println(nums)
	fmt.Println(diagonalSum(nums))
}

func diagonalSum(mat [][]int) int {
	n := len(mat)
	sum := 0
	for i := 0 ; i < n ; i ++ {
		sum += mat[i][i]
		sum += mat[i][n-i-1]
	}
	if n % 2 == 1 {
		sum -= mat[n/2][n/2]
	}

	return sum
}
