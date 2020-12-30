package main

import (
	"container/heap"
	"sort"
)

// 判断可否分成连续的 长度大于或等于3的 子序列
type minHp struct{ sort.IntSlice }

func (h *minHp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *minHp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *minHp) push(v int) { heap.Push(h, v) }
func (h *minHp) pop() int   { return heap.Pop(h).(int) }

func isPossible1(nums []int) bool {
	// 存储各子序列的长度，最后做判断是否每个都大于3
	lens := map[int]*minHp{}

	for _, v := range nums {
		// 初始化
		if lens[v] == nil {
			lens[v] = new(minHp)
		}
		//
		if h := lens[v-1]; h != nil { // 如果有以v-1结尾的子序列
			// 最小堆堆顶，pop出来最小子序列的长度
			prevLen := h.pop()
			// 如果该最小堆已经没有元素了就删掉，因为h!=nil条件不能判断len为，只是判断指针是否为nil
			if h.Len() == 0 {
				delete(lens, v-1)
			}
			// 加上v，长度+1，push进去
			lens[v].push(prevLen + 1)
		} else { // 没有 则新增一个长度为1的以v结尾的子序列
			lens[v].push(1)
		}
	}

	for _, h := range lens {
		if h.IntSlice[0] < 3 {
			return false
		}
	}
	return true
}

// 贪心算法
func isPossible2(nums []int) bool {
	left := make(map[int]int) // 每个数字的剩余个数
	for _, v := range nums {
		left[v]++
	}
	endCnt := make(map[int]int) // 以某个数字结尾的连续子序列的个数
	for _, v := range nums {    // 如果剩余次数为0，跳过
		if left[v] == 0 {
			continue
		}
		if endCnt[v-1] > 0 { // 若存在以 v-1 结尾的连续子序列，则将 v 加到其末尾
			left[v]--
			endCnt[v-1]--
			endCnt[v]++
		} else if left[v+1] > 0 && left[v+2] > 0 { // 否则，生成一个长度为 3 的连续子序列
			left[v]--
			left[v+1]--
			left[v+2]--
			endCnt[v+2]++
		} else { // 无法生成
			return false
		}
	}
	return true
}

func isPossible3(nums []int) bool {
	dp1, dp2, dp3 := 0, 0, 0
	i := 0
	for i < len(nums) {
		start := i
		x := nums[i]
		// 寻找相等的值（跳过相等值
		for i < len(nums) && nums[i] == x {
			i++
		}

		cnt := i - start

		// 如果x不为前面的下一个连续数
		if start > 0 && x != nums[start-1]+1 {
			if dp1+dp2 > 0 {
				return false
			} else {
				dp1 = cnt
				dp2, dp3 = 0, 0
			}
		} else {
			if dp1+dp2 > cnt {
				return false
			}

			left := cnt - dp1 - dp2

			use3 := left
			if left > dp3 {
				use3 = dp3
			}

			dp3 = dp2 + use3
			dp2 = dp1
			dp1 = left - use3

		}

	}
	return dp1 == 0 && dp2 == 0

}
