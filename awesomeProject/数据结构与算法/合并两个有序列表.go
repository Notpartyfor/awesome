package main

import (
	"fmt"
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	res := make([]string, 0)
	for l != nil {
		res = append(res, strconv.Itoa(l.Val))
		l = l.Next
	}
	return strings.Join(res, "→")
}
func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 5,
				Next: &ListNode{
					Val: 7,
					Next: &ListNode{
						Val:  11,
						Next: nil,
					},
				},
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val:  8,
						Next: nil,
					},
				},
			},
		},
	}

	fmt.Println(mergeTwoLists(l1, l2))
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	prehead := &ListNode{}
	result := prehead
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			prehead.Next = l1
			l1 = l1.Next
		} else {
			prehead.Next = l2
			l2 = l2.Next
		}

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
