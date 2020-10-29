package main

import "fmt"

//给定一个字符串 s 和一个非空字符串 p，找到 s 中所有是 p 的字母异位词的子串，返回这些子串的起始索引。
//字母异位词指字母相同，但排列不同的字符串。
//输入:s: "cbaebabacd" p: "abc"
//输出:[0, 6]
//解释:
//起始索引等于 0 的子串是 "cba", 它是 "abc" 的字母异位词。
//起始索引等于 6 的子串是 "bac", 它是 "abc" 的字母异位词
func main() {
	source := "cbaebabacd"
	target := "abc"
	fmt.Println(findAnagrams1(source, target))
}

func findAnagrams1(source, target string) []int {
	res := make([]int, 0)
	sm := make(map[string]int)
	pm := make(map[string]int)
	l1 := len(source)
	l2 := len(target)

	for _, v := range target {
		pm[string(v)] += 1
	}

	for _, v := range source[:l2] {
		sm[string(v)] += 1
	}

	for left, right := 0, l2; right < l1; right++ {
		if isSame(sm, pm) {
			res = append(res, left)
		}
		sm[string(source[left])] -= 1
		sm[string(source[right])] += 1
		left += 1
	}
	return res
}

func isSame(sm, pm map[string]int) bool {
	for k, v := range sm {
		if pm[k] != v {
			return false
		}
	}
	return true
}
