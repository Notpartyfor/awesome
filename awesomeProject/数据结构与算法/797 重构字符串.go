package main

import (
	"container/heap"
	"fmt"
	"sort"
)

// 给定一个字符串S，检查是否能重新排布其中的字母，使得两相邻的字符不同。
//
//若可行，输出任意可行的结果。若不可行，返回空字符串。
//
//示例 1:
//输入: S = "aab"
//输出: "aba"

//示例 2:
//输入: S = "aaab"
//输出: ""

//注意:
//S 只包含小写字母并且长度在[1, 500]区间内

// 最大堆
type hp struct{ sort.IntSlice }

// 要先定义这个变量，Less中需要使用
var count [26]int

// 实现堆的
func (h hp) Less(i, j int) bool  { return count[h.IntSlice[i]] > count[h.IntSlice[j]] } // 要用>，因为用的是最大堆
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{}   { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp) push(v int)         { heap.Push(h, v) }
func (h *hp) pop() int           { return heap.Pop(h).(int) }

func main() {
	fmt.Println(reorganizeString("aab"))
}
func reorganizeString(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}

	count = [26]int{}
	maxCount := 0
	for _, v := range s {
		count[v-'a'] += 1
		if maxCount < count[v-'a'] {
			maxCount = count[v-'a']
		}
	}

	if maxCount > (n+1)/2 {
		return ""
	}

	h := &hp{}
	for i, v := range count {
		if v > 0 {
			h.IntSlice = append(h.IntSlice, i)
		}
	}
	heap.Init(h)

	ans := make([]byte, 0, n)
	// 如果还存在两个或以上
	for len(h.IntSlice) > 1 {
		i, j := h.pop(), h.pop()
		ans = append(ans, byte('a'+i), byte('a'+j))
		if count[i]--; count[i] > 0 {
			h.push(i)
		}
		if count[j]--; count[j] > 0 {
			h.push(j)
		}

	}
	// 如果还存在最后一个，加到最后
	if len(h.IntSlice) > 0 {
		ans = append(ans, byte('a'+h.IntSlice[0]))
	}

	return string(ans)
}
