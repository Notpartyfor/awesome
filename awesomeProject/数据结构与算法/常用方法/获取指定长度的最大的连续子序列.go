package main

import "fmt"

func main() {
	arr := []int{9, 6, 2, 11, 7, 8, 1, 5, 3}
	for i := range arr {
		fmt.Println(getKMost(arr, i))
		fmt.Println(maxSubsequence(arr, i))
	}
}
func getKMost(arr []int, k int) (ret []int) {
	for i, v := range arr {
		// len(ret) 是当前栈内元素个数
		// len(arr)-i 是遍历到现在，数组中还未入栈的个数

		// 表示此时(栈内元素 + 还未入栈的元素) - 1的数量要 >= k (将栈顶替换成当前值 会吃掉一个数)
		// for循环去删除栈顶元素是为了让后面更大的放到栈底
		for len(ret) > 0 && len(ret)+len(arr)-1-i >= k && v > ret[len(ret)-1] {
			fmt.Printf("k = %d, getKMost第%d轮，此时当前值为%d,栈内数据为%v,条件为%d\n", k, i, v, ret, len(ret)+len(arr)-1-i)
			ret = ret[:len(ret)-1]
		}
		// 先入满到k个
		if len(ret) < k {
			ret = append(ret, v)
		}
	}

	return
}

// 返回长度为k的最大子序列
func maxSubsequence(nums []int, k int) []int {
	n := len(nums) - k
	stack := make([]int, 0)
	// 下面是直接套用 从数组中取出n个元素，剩下的数组为最大的函数，所以在之前将n := len(nums) - k
	for _, v := range nums {
		for len(stack) > 0 && n > 0 && v > stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			n--
		}
		stack = append(stack, v)
	}

	stack = stack[:k]
	return stack
}
