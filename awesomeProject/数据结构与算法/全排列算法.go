package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}
func permute(nums []int) [][]int {
	res := make([][]int, 0)
	temp := make([]int, 0)
	// 注意指针
	dfs(nums, temp, &res)
	return res
}
func dfs(nums, temp []int, res *[][]int) {
	fmt.Printf("%v,%v\n", nums, temp)
	// 如果长度一样，就证明取完了，加到res里面
	if len(temp) == len(nums) {
		*res = append(*res, temp)
	} else { // 不然的话就要继续向下取
		for _, n := range nums {
			// 向下取没包含的
			if !contains(temp, n) {
				// 加到temp里面
				temp = append(temp, n)
				// 加了之后再向下取
				dfs(nums, temp, res)
				// 这里精髓
				temp = temp[:len(temp)-1]
			}
		}
	}
}

func contains(source []int, target int) bool {
	for _, v := range source {
		if v == target {
			return true
		}
	}
	return false
}
