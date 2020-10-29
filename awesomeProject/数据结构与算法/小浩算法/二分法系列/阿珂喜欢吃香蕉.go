package main

import (
	"awesomeProject/common"
	"fmt"
)

// 这里总共有 N 堆香蕉，第 i 堆中有piles[i] 根香蕉。
// 警卫已经离开了，将在 H 小时后回来。
// 阿珂可以决定她吃香蕉的速度 K （单位：根/小时），每个小时，她将会选择一堆香蕉，从中吃掉 K 根。

// 如果这堆香蕉少于 K 根，她将吃掉这堆的所有香蕉，然后这一小时内不会再吃更多的香蕉。
// 珂珂喜欢慢慢吃，但仍然想在警卫回来前吃掉所有的香蕉。返回她可以在 H 小时内吃掉所有香蕉的最小速度 K（K 为整数）。

// 输入: piles = [3,6,7,11], H = 8
// 输出: 4

// 输入: piles = [30,11,23,4,20], H = 5
// 输出: 30

// 输入: piles = [30,11,23,4,20], H = 6
// 输出: 23
func main() {
	piles := []int{30, 11, 23, 4, 20}
	H := 5
	fmt.Println(minEatingSpeed(piles, H))
}

// 二分法介绍
func binarySearch(nums []int, des int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if des == nums[mid] {
			return mid
		} else if des < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func minEatingSpeed(piles []int, H int) int {
	left := 1 // 假设要吃的最小速度为1
	right := common.MaxInSlice(piles)

	for left < right {
		mid := (left + right) / 2
		if canEat(piles, mid, H) { // 所用时间过长，要加速
			left = mid + 1
		} else {
			right = mid
		}

	}
	return left

}

func canEat(piles []int, speed, H int) bool {
	sum := 0
	for _, v := range piles {
		sum += (v + speed - 1) / speed
		// 向上取整小技巧：加了一个H，其实就相当于商加了1，所以加1后再除以H（向下取整）就等于原先的向上取整了

	}
	return sum > H
}
