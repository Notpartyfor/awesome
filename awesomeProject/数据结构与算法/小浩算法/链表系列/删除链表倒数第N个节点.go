package main

import (
	"fmt"
)

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	fmt.Println(head)
	fmt.Println(removeNthFromEnd(head, 2))
}

func removeNthFromEnd(head *ListNode, N int) *ListNode {
	// 定义一个烧饼节点
	result := &ListNode{}
	result.Next = head
	cur := result
	var pre *ListNode
	i := 1

	for head != nil {
		if i > N-1 {
			pre = cur
			cur = cur.Next
		}
		head = head.Next
		i++
	}
	pre.Next = pre.Next.Next
	return result.Next
}
