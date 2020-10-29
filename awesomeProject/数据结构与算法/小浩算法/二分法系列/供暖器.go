package main

import (
	"awesomeProject/common"
	"fmt"
	"sort"
)

// 冬季已经来临。你的任务是设计一个有固定加热半径的供暖器向所有房屋供暖。
// 现在，给出位于一条水平线上的房屋和供暖器的位置，找到可以覆盖所有房屋的最小加热半径。
// 所以，你的输入将会是房屋和供暖器的位置。你将输出供暖器的最小加热半径

// 输入: [1,2,3,4],[1,4]
// 输出: 1
// 解释: 在位置1, 4上有两个供暖器。我们需要将加热半径设为1，这样所有房屋就都能得到供暖。
func main() {
	houses := []int{1, 2, 3, 4, 5, 6, 7, 8}
	heaters := []int{4, 9}
	fmt.Println(findRadius(houses, heaters))
}

func findRadius(houses, heaters []int) int {
	res := 0
	n := len(heaters)
	sort.Ints(heaters)
	for _, house := range houses {
		// 二分法找不小于house的第一个值
		// 其实也就是找到house后面一个暖气站的下标
		left, right := 0, n
		//
		for left < right {
			mid := (left + right) / 2
			if house > heaters[mid] {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 前面和后面暖气的距离，取最小
		dist1, dist2 := 0, 0
		if right == 0 {
			dist1 = common.IntMax
		} else {
			dist1 = common.Abs(house - heaters[right-1])
		}
		if right == n {
			dist2 = common.IntMax
		} else {
			dist2 = common.Abs(house - heaters[right])
		}

		res = common.Max(res, common.Min(dist1, dist2))

	}
	return res
}
