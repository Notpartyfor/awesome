package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key, val  int
	pre, next *Node
}

type LRU struct {
	m          map[int]*Node
	lock       sync.Mutex
	cap        int
	head, tail *Node
}

func main() {
	LRU := NewLRUMine(2)
	LRU.Put(1, 11)
	LRU.Put(2, 22)
	LRU.Put(3, 33)
	fmt.Println(LRU.Get(1))
	fmt.Println(LRU.Get(2))
	fmt.Println(LRU.Get(3))
}

func NewLRUMine(cap int) *LRU {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.pre = head
	m := make(map[int]*Node)
	return &LRU{
		m:    m,
		cap:  cap,
		head: head,
		tail: tail,
	}

}

func (l *LRU) Get(key int) int {
	// defer l.lock.Unlock() 会损失掉一部分性能
	// 代价就是详细点写
	// 缓存击中
	l.lock.Lock()
	if n, ok := l.m[key]; ok {
		// 将他放到链表首位
		l.MoveToHead(n)
		return n.val
	} else {
		return -1
	}
}

func (l *LRU) Put(key, val int) {
	// 缓存击中
	if n, ok := l.m[key]; ok {
		n.val = val
		l.MoveToHead(n)
	} else {
		node := &Node{
			key:  key,
			val:  val,
			pre:  nil,
			next: nil,
		}
		// 未击中则需要增加新的节点
		// 还有剩余长度时直接增加，没有的时候删去最近最不常用的点，也就是链表末位
		if len(l.m) == l.cap {
			// 这里要注意要先删除缓存再删除节点，不然指向会出错哦
			delete(l.m, l.tail.pre.key)
			l.RemoveNode(l.tail.pre)
		}

		l.AddNode(node)
		l.m[key] = node
	}
}

func (l *LRU) MoveToHead(node *Node) {
	// 加到首位
	l.AddNode(node)
	// 将原本的删除
	l.RemoveNode(node)
}

func (l *LRU) AddNode(node *Node) {
	// 处理首位
	node.next = l.head.next
	l.head.next.pre = node
	// 处理头部的哨兵节点

	node.pre = l.head
	l.head.next = node
}

func (l *LRU) RemoveNode(node *Node) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
