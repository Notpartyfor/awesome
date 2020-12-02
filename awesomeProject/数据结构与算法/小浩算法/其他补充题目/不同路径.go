package main

import "fmt"

// 一个机器人位于一个 m x n 网格的左上角，起始点在下图中标记为“Start”。
// 机器人每次只能向下或者向右移动一步。机器人试图达到网格的右下角，在下图中标记为“Finish”。
//
// 问：总共有多少条不同的路径？
func main() {
	fmt.Println(path(7, 3))
	fmt.Println(pathUpgrade(7, 3))
	fmt.Println(C(7, 3))
}

// 一般做法
func path(m, n int) int {
	// 初始化dp
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 如果在第一列，那到达的方式只有一种
	for i := 0; i < m; i++ {
		dp[i][0] = 1
	}
	// 第一行
	for i := 0; i < n; i++ {
		dp[0][i] = 1
	}
	// 这里的dp[i][j] 表示到达那里有多少种方法
	// 如果不在第一行或者第一列，那dp[i][j] = dp[i-1][j] + dp[i][j-1] 一直累计到dp[m-1][n-1]就完事了
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

// 优化做法
// 我们定义状态dp[j]，用来表示当前行到达第j列的最多路径
// 此时dp[j] = dp[j-1] + dp[j]
// dp[j-1]已经更新为当前行到第j-1列，即左边的路径数，而第一列的其实没有左边的，所以一开始dp[0]只是由上面赋值
// 而等式右边的dp[j]还未更新，表示上边的路径数
// 加起来便是新的路径数

func pathUpgrade(m, n int) int {
	dp := make([]int, n)
	// 第一行都为1
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	// 然后每一行都走下去
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] = dp[j-1] + dp[j]
		}
	}
	return dp[n-1]
}

// 试试公式法，好像不太对
func C(m, n int) int {
	res := 1
	resm := m
	resn := 1
	for n > 0 {
		res *= m
		m -= 1
		resn *= n
		n -= 1
	}
	return (res / resn) - resm
}
