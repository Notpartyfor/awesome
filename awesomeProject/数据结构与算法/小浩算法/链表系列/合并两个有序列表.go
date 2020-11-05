package main

import (
	"awesomeProject/persistance"
	"fmt"
)

func main() {
	l1 := &persistance.ListNode{
		Val: 1,
		Next: &persistance.ListNode{
			Val: 3,
			Next: &persistance.ListNode{
				Val: 5,
				Next: &persistance.ListNode{
					Val: 7,
					Next: &persistance.ListNode{
						Val:  11,
						Next: nil,
					},
				},
			},
		},
	}
	l2 := &persistance.ListNode{
		Val: 1,
		Next: &persistance.ListNode{
			Val: 3,
			Next: &persistance.ListNode{
				Val: 4,
				Next: &persistance.ListNode{
					Val: 5,
					Next: &persistance.ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *persistance.ListNode, l2 *persistance.ListNode) *persistance.ListNode {
	// 定义一个烧饼节点
	prehead := &persistance.ListNode{}
	// 初始化返回节点
	result := prehead

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			// 指向小的
			prehead.Next = l1
			// 推进l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			// 推进l2
			l2 = l2.Next
		}
		// 推进prehead
		prehead = prehead.Next
	}
	// 到这里就有一个终止了，此时接上另一个就好了
	if l1 != nil {
		prehead.Next = l1
	}

	if l2 != nil {
		prehead.Next = l2
	}
	return result.Next
}
