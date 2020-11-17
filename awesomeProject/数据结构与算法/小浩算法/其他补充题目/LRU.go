package main

type LinkNode struct {
	key, val  int
	pre, next *LinkNode
}

type LRUCache struct {
	m          map[int]*LinkNode
	cap        int
	head, tail *LinkNode
}

func NewLRUCache(capacity int) LRUCache {
	// 初始化头尾
	head := &LinkNode{0, 0, nil, nil}
	tail := &LinkNode{0, 0, nil, nil}
	head.next = tail
	tail.pre = head

	return LRUCache{
		m:    make(map[int]*LinkNode),
		cap:  capacity,
		head: head,
		tail: tail,
	}
}

func (l *LRUCache) Get(key int) int {
	cache := l.m
	if node, ok := cache[key]; ok {
		l.MoveToHead(node)
		return node.val
	} else {
		return -1
	}
}

func (l *LRUCache) PUT(key, val int) {
	tail := l.tail
	cache := l.m

	if node, ok := cache[key]; ok {
		node.val = val
		l.MoveToHead(node)
	} else {
		v := &LinkNode{
			key:  key,
			val:  val,
			pre:  nil,
			next: nil,
		}
		// 如果满了，需要删掉最少使用的节点
		if len(cache) == l.cap {
			delete(cache, tail.key)
			l.DeleteNode(tail)
		}

		l.AddNode(v)
		cache[key] = v
	}
}

func (l *LRUCache) MoveToHead(node *LinkNode) {
	l.AddNode(node)
	l.DeleteNode(node)
}
func (l *LRUCache) AddNode(node *LinkNode) {
	head := l.head

	node.next = head.next
	head.next.pre = node

	node.pre = head
	head.next = node
}
func (l *LRUCache) DeleteNode(node *LinkNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
}
