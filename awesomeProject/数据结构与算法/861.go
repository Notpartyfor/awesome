package main

import (
	"fmt"
)

func main() {
	A := [][]int{
		{0, 0, 1, 1},
		{1, 0, 1, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(matrixScore(A))
}

// 实际上，反转的顺序其实并不影响结果..
// 我们不妨全部用行的反转，再用列的反转

// 肯定优先考虑第一列，即尽量保证第一行全部为1
// 而因为后面可以用列的反转将全0变为全1
// 所以要考虑的就有两种情况，在第一列全0或者全1下操作
// 实际上最后因为全0和全1而不同的反转列数是 互补的

func matrixScore(A [][]int) int {
	res := 0
	l1 := len(A)
	l2 := len(A[0])
	// 先进行 行变换，将其第一列全变为1
	for _, a := range A {
		if a[0] == 0 {
			reverseRow(a)
		}
	}
	// 先把第一列加上呗
	res += l1 * (1 << (l2 - 1))
	// 将剩下的根据0的个数多寡来决定是否进行列变换
	for i := 1; i < l2; i++ {
		sum := 0
		for _, a := range A {
			sum += a[i]
		}
		// 如果0太多，就反转
		// 事实上这一步可以省略反转的实际操作，因为只需要获取分数
		if sum <= l1/2 {
			//for _, a := range A {
			//	a[i] = reverse(a[i])
			//}
			res += (l1 - sum) * (1 << (l2 - i - 1))
		} else {
			res += sum * (1 << (l2 - i - 1))
		}
	}
	return res
}

func reverseRow(a []int) []int {
	for i := range a {
		a[i] = reverse(a[i])
	}
	return a
}

func reverse(a int) int {
	if a == 0 {
		return 1
	}
	return 0
}
