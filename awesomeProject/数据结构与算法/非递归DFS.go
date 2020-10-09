package main

import "awesomeProject/persistance"

func main() {

}

func DFS(root persistance.TreeNode) []int {
	res := make([]int, 0)
	Stack := persistance.NewSliceEntry()
	Stack.Push(root)
	for !Stack.IsEmpty() {
		var node = Stack.Top().(persistance.TreeNode)
		Stack.Pop()
		_ = node
	}
	return res
}
