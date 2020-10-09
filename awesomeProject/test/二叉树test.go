package main

import (
	"awesomeProject/persistance"
	"fmt"
)

func main() {
	root := &persistance.TreeNode{
		Val: 5,
		Left: &persistance.TreeNode{
			Val: 3,
			Left: &persistance.TreeNode{
				Val:   2,
				Left:  nil,
				Right: nil,
			},
			Right: &persistance.TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &persistance.TreeNode{
			Val: 7,
			Left: &persistance.TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: &persistance.TreeNode{
				Val:   9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	fmt.Println(root)
	//fmt.Println(root.IsBST())
	//fmt.Println(root.SearchBST(2))
	fmt.Println(root.DeleteNode(3))

}

