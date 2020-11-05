package main

import (
	"awesomeProject/persistance"
	"fmt"
)

func main() {
	head := &persistance.ListNode{
		Val: 1,
		Next: &persistance.ListNode{
			Val: 2,
			Next: &persistance.ListNode{
				Val: 3,
				Next: &persistance.ListNode{
					Val: 4,
					Next: &persistance.ListNode{
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

func removeNthFromEnd(head *persistance.ListNode, N int) *persistance.ListNode {
	// 定义一个烧饼节点
	result := &persistance.ListNode{}
	result.Next = head

	cur := result
	pre := &persistance.ListNode{}
	i := 1

	for head != nil {
		if i > N-1 {
			// head指针当移动到对的位置（通过i判断），开始移动cur
			pre = cur
			cur = cur.Next
		}
		// head指针移动 代表 进度条
		head = head.Next
		i++
	}
	// 删除 pre的下一个节点，也就是cur节点
	pre.Next = pre.Next.Next

	// 只是通过烧饼节点来返回整条链表
	return result.Next
}
