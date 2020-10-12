package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 一副扑克牌不算大小王刚好是52张，随意洗牌。
// 如果出现连续三张牌，花色依次是红黑黑，那么玩家A加一分；
// 同时把翻开了的牌都丢掉，继续一张张翻没翻开的牌；
// 类似地，一 旦出现连续三张牌恰好是黑黑红，则玩家B得一分，弃掉已翻开的牌后继续。结果会如何呢？

const (
	RED   = 0
	BLACK = 1
)

func main() {
	nums := NewPokers()
	fmt.Println(nums)
	fmt.Println(GetScore(nums))
}

func NewPokers() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	nums := make([]int, 0)
	for i := 0; i < 26; i++ {
		nums = append(nums, 0)
		nums = append(nums, 1)
	}
	for i := 0; i < 52; i++ {
		num := r.Intn(i + 1)
		nums[i], nums[num] = nums[num], nums[i]
	}
	//fmt.Println(nums)
	//time.Sleep(100)
	for i := 0; i < 51; i++ {
		num := r.Intn(i + 2)
		nums[i], nums[num] = nums[num], nums[i]
	}
	//fmt.Println(nums)
	//time.Sleep(100)
	for i := 0; i < 50; i++ {
		num := r.Intn(i + 3)
		nums[i], nums[num] = nums[num], nums[i]
	}

	return nums
}

func GetScore(nums []int) (a, b []int) {
	A := 0
	B := 0
	for i := 0; i < 50; i++ {
		//
		if nums[i] == RED && nums[i+1] == BLACK && nums[i+2] == BLACK {
			a = append(a, i)
			A += 1
			i += 2
		} else if nums[i] == BLACK && nums[i+1] == BLACK && nums[i+2] == RED {
			b = append(b, i)
			B += 1
			i += 2
		}
	}
	return a, b
}
